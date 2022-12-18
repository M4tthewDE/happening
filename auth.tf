data "archive_file" "auth_zip" {
  type        = "zip"
  source_file = "bin/auth"
  output_path = "bin/auth.zip"
}

resource "aws_lambda_function" "auth_lambda" {
  filename         = data.archive_file.auth_zip.output_path
  function_name    = "auth-${local.app_id}"
  handler          = "auth"
  source_code_hash = base64sha256(data.archive_file.auth_zip.output_path)
  runtime          = "go1.x"
  role             = aws_iam_role.lambda_exec.arn
}
