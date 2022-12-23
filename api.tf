resource "random_password" "eventsub_secret" {
  length  = 32
  special = false
}

# Assume role setup
resource "aws_iam_role" "lambda_exec" {
  name_prefix = local.app_id

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_role_policy" "dynamodb" {
  name = "dynamodb"
  role = aws_iam_role.lambda_exec.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "dynamodb:*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

# Attach role to Managed Policy
variable "iam_policy_arn" {
  description = "IAM Policy to be attached to role"
  type        = list(string)

  default = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]
}

resource "aws_iam_policy_attachment" "role_attach" {
  name       = "policy-${local.app_id}"
  roles      = [aws_iam_role.lambda_exec.id]
  count      = length(var.iam_policy_arn)
  policy_arn = element(var.iam_policy_arn, count.index)
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "bin/api"
  output_path = "bin/api.zip"
}

resource "aws_lambda_function" "lambda_func" {
  filename         = data.archive_file.lambda_zip.output_path
  function_name    = "api-${local.app_id}"
  handler          = "api"
  source_code_hash = base64sha256(data.archive_file.lambda_zip.output_path)
  runtime          = "go1.x"
  role             = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      EVENTSUB_SECRET  = random_password.eventsub_secret.result
      TWITCH_CLIENT_ID = var.TWITCH_CLIENT_ID
      API_URL          = var.API_URL
    }
  }
}
