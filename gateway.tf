resource "aws_api_gateway_rest_api" "api" {
  name = local.app_id
}

resource "aws_api_gateway_resource" "proxy" {
  path_part   = "{proxy+}"
  parent_id   = aws_api_gateway_rest_api.api.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.api.id
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = aws_api_gateway_rest_api.api.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "proxy_root" {
  rest_api_id   = aws_api_gateway_rest_api.api.id
  resource_id   = aws_api_gateway_rest_api.api.root_resource_id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_method.method.resource_id
  http_method             = aws_api_gateway_method.method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda_func.invoke_arn
}

resource "aws_api_gateway_integration" "integration_root" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_method.proxy_root.resource_id
  http_method             = aws_api_gateway_method.proxy_root.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda_func.invoke_arn
}

resource "aws_api_gateway_deployment" "api_deployment" {
  depends_on = [
    aws_api_gateway_integration.integration,
    aws_api_gateway_integration.integration_root,
  ]

  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name  = "api"
}

resource "aws_lambda_permission" "lambda_permission" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_func.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_deployment.api_deployment.execution_arn}/*/*"
}

# needs to be manually verified at cloudflare
resource "aws_acm_certificate" "cert" {
  domain_name       = "happening.fdm.com.de"
  validation_method = "DNS"

  tags = {
    Environment = "develop"
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "happening" {
  certificate_arn = aws_acm_certificate.cert.arn
}

resource "aws_api_gateway_domain_name" "happening" {
  domain_name              = "happening.fdm.com.de"
  regional_certificate_arn = aws_acm_certificate_validation.happening.certificate_arn

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}
