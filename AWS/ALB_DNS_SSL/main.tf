resource "aws_key_pair" "danikey" {
  key_name   = "danikey"
  public_key = file("/home/daniel/.ssh/id_rsa.pub")
}

resource "aws_instance" "daniinstance" {
  key_name        = aws_key_pair.danikey.key_name
  ami             = "ami-099a8245f5daa82bf"
  instance_type   = "t2.micro"
  security_groups = [aws_security_group.allow_ssh.name]
  connection {
    type        = "ssh"
    user        = "ec2-user"
    private_key = file("/home/daniel/.ssh/id_rsa")
    host        = self.public_ip
  }
  provisioner "remote-exec" {
    inline = [
      "sudo amazon-linux-extras install nginx1.12 -y && sudo service nginx start"
    ]
  }
}

resource "aws_security_group" "allow_ssh" {
  name        = "allow_ssh"
  description = "Allow SSH inbound traffic"
  #vpc_id      = aws_vpc.main.id #Implying default, YOLO
  ingress {
    description = "SSH from Internet"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "allow_ssh"
  }
}

resource "aws_security_group" "alb_allow_https" {
  name        = "allow_https"
  description = "Allow https inbound traffic"
  #vpc_id      = aws_vpc.main.id #Implying default, YOLO
  ingress {
    description = "https from Internet"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "allow_https"
  }
}

resource "aws_security_group" "http_only_from_lb" {
  name        = "http_only_from_lb"
  description = "Allow HTTP traffic only from load balancer"
  #vpc_id      = aws_vpc.main.id #Implying default, YOLO
  ingress {
    description = "HTTP from load balancer"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    security_groups = [aws_security_group.alb_allow_https.id]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "http_only_from_lb"
  }
}

resource "aws_lb_target_group" "danitg" {
  name     = "danitg"
  port     = 80
  protocol = "HTTP"
  vpc_id   = data.aws_vpc.current_vpc.id
}

resource "aws_lb_target_group_attachment" "danitg_attachment" {
  target_group_arn = aws_lb_target_group.danitg.arn
  target_id        = aws_instance.daniinstance.id
  port             = 80
}

resource "aws_lb" "dani_lb" {
  name                       = "dani-lb"
  internal                   = false
  load_balancer_type         = "application"
  security_groups            = [aws_security_group.alb_allow_https.id]
  subnets                    = data.aws_subnet_ids.current_subnets.ids
  enable_deletion_protection = false

  tags = {
    Environment = "experimental"
  }
}

resource "aws_lb_listener" "dani_lb_listener_https" {
  load_balancer_arn = aws_lb.dani_lb.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = aws_acm_certificate.wildcard-cert.arn
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.danitg.arn
  }
}

resource "aws_route53_record" "alb_wildcard_dns" {
  zone_id = data.terraform_remote_state.dcalvo-dev-zone.outputs.zone_id
  name    = "*.plswork.dcalvo.dev"
  type    = "CNAME"
  ttl     = "300"
  records = [aws_lb.dani_lb.dns_name]
}

resource "aws_acm_certificate" "wildcard-cert" {
  domain_name       = "*.plswork.dcalvo.dev"
  validation_method = "DNS"

  tags = {
    Environment = "plswork.dcalvo.dev"
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "acm-validation-wildcard-cert" {
  for_each = {
    for dvo in aws_acm_certificate.wildcard-cert.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = data.terraform_remote_state.dcalvo-dev-zone.outputs.zone_id
}

resource "aws_acm_certificate_validation" "acm-validation-wildcard-cert" {
  certificate_arn         = aws_acm_certificate.wildcard-cert.arn
  validation_record_fqdns = [for record in aws_route53_record.acm-validation-wildcard-cert : record.fqdn]
}