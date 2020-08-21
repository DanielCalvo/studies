resource "aws_launch_template" "foo" {
  name = "foo"
  instance_type = "t2.micro"
  image_id = "ami-07d9160fa81ccffb5"
  key_name = "dani"
  user_data = "IyEvYmluL2Jhc2gKeXVtIHVwZGF0ZSAteQp5dW0gaW5zdGFsbCAteSBodHRwZApzeXN0ZW1jdGwgc3RhcnQgaHR0cGQuc2VydmljZQpzeXN0ZW1jdGwgZW5hYmxlIGh0dHBkLnNlcnZpY2UKRUMyX0FWQUlMX1pPTkU9JChjdXJsIC1zIGh0dHA6Ly8xNjkuMjU0LjE2OS4yNTQvbGF0ZXN0L21ldGEtZGF0YS9wbGFjZW1lbnQvYXZhaWxhYmlsaXR5LXpvbmUpCmVjaG8gIjxoMT5IZWxsbyBXb3JsZCBmcm9tICQoaG9zdG5hbWUgLWYpIGluIEFaICRFQzJfQVZBSUxfWk9ORSA8L2gxPiIgPiAvdmFyL3d3dy9odG1sL2luZGV4Lmh0bW="

}

resource "aws_autoscaling_group" "bar" {
  availability_zones = ["eu-west-1a"]
  desired_capacity   = 2
  max_size           = 3
  min_size           = 2
  health_check_grace_period = 120

  health_check_type    = "ELB"
  load_balancers= ["${aws_elb.web_elb.id}" ]

  launch_template {
    id      = aws_launch_template.foo.id
    version = "$Latest"
  }
}