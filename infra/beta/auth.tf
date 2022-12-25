data "archive_file" "auth_zip" {
  type        = "zip"
  source_file = "../../bin/auth"
  output_path = "../../bin/auth.zip"
}

resource "aws_lambda_function" "auth_lambda" {
  filename         = data.archive_file.auth_zip.output_path
  function_name    = "auth-${local.app_id}"
  handler          = "auth"
  source_code_hash = base64sha256(data.archive_file.auth_zip.output_path)
  runtime          = "go1.x"
  role             = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      TWITCH_SECRET    = var.TWITCH_SECRET
      TWITCH_CLIENT_ID = var.TWITCH_CLIENT_ID
      TABLE_NAME       = "auth-${local.app_id}"
    }
  }
}

resource "aws_cloudwatch_event_rule" "auth" {
  name                = "auth-event"
  description         = "Refresh twitch auth token stored in dynamodb"
  schedule_expression = "rate(1 minute)"
}


resource "aws_cloudwatch_event_target" "auth_lambda" {
  rule      = aws_cloudwatch_event_rule.auth.name
  target_id = "auth-lambda"
  arn       = aws_lambda_function.auth_lambda.arn
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_auth_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.auth_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.auth.arn
}
resource "aws_dynamodb_table" "auth-table" {
  name         = "auth-${local.app_id}"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"
  range_key    = "access_token"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "access_token"
    type = "S"
  }
}

resource "aws_dynamodb_table" "permissions-table" {
  name         = "permissions-${local.app_id}"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"
  range_key    = "user_id"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "user_id"
    type = "S"
  }
}

resource "aws_dynamodb_table_item" "default_permission" {
  table_name = aws_dynamodb_table.permissions-table.name
  hash_key   = aws_dynamodb_table.permissions-table.hash_key
  range_key  = aws_dynamodb_table.permissions-table.range_key

  item = <<ITEM
  {
    "id": {"S": "${uuid()}"},
    "user_id": {"S": "116672490"},
    "permissions": {"SS": ["ALL"]}
  }
  ITEM
}
