output "output" {
  value = "test"
}

module "test" {
  source = "../terraform-backend-modules-test/test-module"
  name   = "testing"
}

resource "aws_s3_bucket" "b" {
  bucket = "my-tf-test-bucket"


# deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

# deprecated: Use the aws_s3_bucket_logging resource instead.

# A new 'object_lock_enabled' attribute is now available
  object_lock_enabled = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
resource "aws_s3_bucket" "fasdfasd" {
  bucket = "my-tf-test-bucket"


# deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

# deprecated: Use the aws_s3_bucket_logging resource instead.
  logging = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket" "d" {
  bucket = "my-tf-test-bucket"


# deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

# deprecated: Use the aws_s3_bucket_logging resource instead.
  logging = {}

  object_lock_configuration {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket" "c" {
  bucket = "my-tf-test-bucket"


# deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

# deprecated: Use the aws_s3_bucket_logging resource instead.
  logging = {}

# A new 'object_lock_enabled' attribute is now available
  object_lock_enabled = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}


# A new resource named aws_s3_bucket_acl is now available
resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.b.id
  acl    = "private"
}
data "aws_s3_account_public_access_block" "example" {
}


# A new resource named aws_s3_bucket_website_configuration is now available
resource "aws_s3_bucket_website_configuration" "test" {
  bucket = aws_s3_bucket.b.bucket
}


module "module_example_complete" {
  source  = "cloudposse/module/example//examples/complete"
  version = "1.0.0"
  # insert the 15 required variables here
}

resource "google_runtimeconfig_config_iam_binding" "test" {
  config  = ""
  members = []
  role    = ""
}