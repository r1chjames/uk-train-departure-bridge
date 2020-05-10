provider "aws" {
  region      = "eu-west-1"
  max_retries = 20
}

terraform {
  required_version = ">= 0.12"
}

data "aws_caller_identity" "current" {}