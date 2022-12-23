# Input variable definitions

variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-east-1"
}

variable "APP_ENV" {
  description = "Application environment tag"
}

resource "random_id" "unique_suffix" {
  byte_length = 2
}

variable "app_name" {
  description = "Application name"
  default     = "happening"
}

locals {
  app_id = "${lower(var.app_name)}-${lower(var.APP_ENV)}-${random_id.unique_suffix.hex}"
}
