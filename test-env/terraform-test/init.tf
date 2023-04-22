terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.74.3"
    }
  }
}

provider "aws" {
  shared_config_files      = ["/Users/tf_user/.aws/conf"]
  shared_credentials_files = ["/Users/tf_user/.aws/creds"]
  profile                  = "terraform"
}