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
