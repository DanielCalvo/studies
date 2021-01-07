locals {
  image_name = "helloworld01"
}

resource "aws_ecr_repository" "ecr_repository" {
  name                 = local.image_name
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}

#Not a good solution, but let's roll with it for now
resource "null_resource" "docker_login" {
  depends_on = [aws_ecr_repository.ecr_repository]
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin ${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com"
  }
}

#AAaaah my eyes hurt!
resource "null_resource" "docker_build" {
  depends_on = [aws_ecr_repository.ecr_repository, null_resource.docker_login]
  provisioner "local-exec" {
    command = "docker build -t ${local.image_name} . "
  }
}

#Ok now I'm just messing around
resource "null_resource" "docker_tag_and_push" {
  depends_on = [aws_ecr_repository.ecr_repository, null_resource.docker_login, null_resource.docker_build]
  provisioner "local-exec" {
    command = "docker tag ${local.image_name}:latest ${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/${local.image_name}:latest && docker push ${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/${local.image_name}:latest"
  }
}

