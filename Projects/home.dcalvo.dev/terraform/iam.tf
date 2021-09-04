resource "aws_iam_user" "orangepi" {
  name = "orangepi"
  path = "/"
}

resource "aws_iam_access_key" "orangepi" {
  user = aws_iam_user.orangepi.name
}

resource "aws_iam_user_policy_attachment" "orange_policy_attachment" {
  user = aws_iam_user.orangepi.name
  policy_arn = aws_iam_policy.orangepi_route53.arn
}

resource "aws_iam_policy" "orangepi_route53" {
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "route53:*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}
