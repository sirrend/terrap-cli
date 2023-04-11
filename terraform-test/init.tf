terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.74.3"
    }

    google = {
      source = "hashicorp/google"
      version = "3.90.1"
    }
  }
}