resource "aws_elb" "web_elb" {
  name = "web-elb"
  cross_zone_load_balancing   = true

    availability_zones = ["eu-west-1a","eu-west-1b", "eu-west-1c"]

  health_check {
    healthy_threshold = 2
    unhealthy_threshold = 2
    timeout = 3
    interval = 30
    target = "HTTP:80/"
  }
  listener {
    lb_port = 80
    lb_protocol = "http"
    instance_port = "80"
    instance_protocol = "http"
  }
}