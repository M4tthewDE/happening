# Input variable definitions

variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-east-1"
}

resource "random_id" "unique_suffix" {
  byte_length = 2
}

variable "app_name" {
  description = "Application name"
  default     = "happening"
}

variable "app_env" {
  description = "Environment name"
  default     = "beta"
}

locals {
  app_id = "${lower(var.app_name)}-${lower(var.app_env)}-${random_id.unique_suffix.hex}"
}

locals {
  sub_domain = "${var.app_env}-happening"
}

variable "TWITCH_SECRET" {
  type      = string
  sensitive = true
}

variable "TWITCH_CLIENT_ID" {
  type      = string
  sensitive = true
}
