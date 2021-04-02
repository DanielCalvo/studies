module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = var.cluster_name
  cluster_version = "1.19"
  subnets         = module.vpc.private_subnets

  tags = {
    Environment = "var.cluster_name"
  }

  vpc_id = module.vpc.vpc_id

  workers_group_defaults = {
    root_volume_type = "gp2"
  }

  //Adding policies to nodes isn't the best, but it's the fastest way forward
  workers_additional_policies = [
    aws_iam_policy.cert-manager-policy.arn,
    aws_iam_policy.external-dns-policy.arn
  ]

  worker_groups = [
    {
      name                 = "worker-group-4"
      instance_type        = "t3a.medium"
      asg_desired_capacity = 2
    }
  ]
}