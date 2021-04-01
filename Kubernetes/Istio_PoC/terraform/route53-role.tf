resource "aws_iam_role" "route-53-cert-manager-role" {
  name               = "route-53-cert-manager-role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.route-53-cert-manager-role.name
  policy_arn = aws_iam_policy.route-53-cert-manager-policy.arn
}

resource "aws_iam_policy" "route-53-cert-manager-policy" {
  name        = "route-53-cert-manager-policy"
  path        = "/"
  description = "route-53-cert-manager-policy"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : "route53:GetChange",
        "Resource" : "arn:aws:route53:::change/*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "route53:ChangeResourceRecordSets",
          "route53:ListResourceRecordSets"
        ],
        "Resource" : "arn:aws:route53:::hostedzone/*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "route53:ListHostedZonesByName",
          "route53:ListHostedZones",
          "route53:ListResourceRecordSets"
          ],
        "Resource" : "*"
      }
    ]
  })
}

