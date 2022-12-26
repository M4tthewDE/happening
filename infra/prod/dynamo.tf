
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
    "user_id": {"S": "116672490"}
  }
  ITEM
}
