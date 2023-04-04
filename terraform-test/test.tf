resource "aws_s3_bucket" "b" {
  bucket = "my-tf-test-bucket"


# deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

# deprecated: Use the aws_s3_bucket_logging resource instead.
  logging = {}

# An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
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

# An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
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


