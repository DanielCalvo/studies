resource "aws_iam_policy" "worker_policy_dns" {
  name        = "worker-policy-dns"
  description = "Worker policy for DNS management"
  policy      = file("iam-policy-dns.json")
}