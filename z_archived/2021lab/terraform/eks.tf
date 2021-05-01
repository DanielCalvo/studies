data "aws_eks_cluster" "cluster" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.eks.cluster_id
}

resource "aws_iam_policy" "worker_policy_alb" {
  name        = "worker-policy-alb"
  description = "Worker policy for the ALB Ingress"
  policy      = file("iam-policy-alb.json")
}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = local.cluster_name
  cluster_version = "1.17"
  subnets         = module.vpc.private_subnets

  tags = { #TODO: Change these
    Environment = "training"
    GithubRepo  = "terraform-aws-eks"
    GithubOrg   = "terraform-aws-modules"
  }

  vpc_id                      = module.vpc.vpc_id
  #workers_additional_policies = [aws_iam_policy.worker_policy_alb.arn, aws_iam_policy.worker_policy_dns.arn]
  workers_additional_policies = [aws_iam_policy.worker_policy_alb.arn]
  #Hey is this DNS stuff still relevant now that we're no longer doing this?
  #TODO: The dns policy above is insecure and not recommended. I'm speedrunning this setup, but this should be adjusted later as described here: https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/aws.md

  worker_groups = [
    {
      name                          = "worker-group-1"
      instance_type                 = "t3.micro"
      asg_desired_capacity          = 3
      additional_security_group_ids = [aws_security_group.all_worker_mgmt.id]
    },
  ]
}