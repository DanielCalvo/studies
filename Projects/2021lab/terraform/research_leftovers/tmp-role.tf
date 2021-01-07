locals {
  cluser_oidc_issuer = trimprefix("https://oidc.eks.eu-west-1.amazonaws.com/id/A19F675117AD4C348A0158AFA1DCF4AA", "https://")
}

resource "aws_iam_role" "k8s_dns_role" {
  name = "test_role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::${data.aws_caller_identity.current.account_id}:oidc-provider/${local.cluser_oidc_issuer}"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "${local.cluser_oidc_issuer}:sub": "system:serviceaccount:default:external-dns"
        }
      }
    }
  ]
}
EOF
  tags = {
    tag-key = "tag-value"
  }
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.k8s_dns_role.name
  policy_arn = aws_iam_policy.worker_policy_dns.arn
}

# Don't forget to attach the policy (with the policy ARN) to the role!
# You need to create a service account now

#Make sure to use variables on the "${module.eks.cluster_oidc_issuer_url}:sub": "system:serviceaccount:default:external-dns" line