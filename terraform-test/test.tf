resource "aws_s3_bucket" "b" {
  bucket = "my-tf-test-bucket"


# The acceleration_status attribute under the aws_s3_bucket resource is now deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

# The logging block_type under the aws_s3_bucket resource is now deprecated: Use the aws_s3_bucket_logging resource instead.
  logging = {}




# The required parameter under the aws_s3_bucket resource was removed: .
# The object_lock_enabled attribute under the aws_s3_bucket resource is now optional.
# The object_lock_enabled attribute under the aws_s3_bucket resource is now deprecated: Use the top-level parameter object_lock_enabled instead.
# An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
  object_lock_enabled = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
resource "aws_s3_bucket" "fasdfasd" {
  bucket = "my-tf-test-bucket"


  # The acceleration_status attribute under the aws_s3_bucket resource is now deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

  # The logging block_type under the aws_s3_bucket resource is now deprecated: Use the aws_s3_bucket_logging resource instead.
  logging = {}




  # The required parameter under the aws_s3_bucket resource was removed: .
  # The object_lock_enabled attribute under the aws_s3_bucket resource is now optional.
  # The object_lock_enabled attribute under the aws_s3_bucket resource is now deprecated: Use the top-level parameter object_lock_enabled instead.
  # An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
  object_lock_enabled = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket" "d" {
  bucket = "my-tf-test-bucket"


  # The acceleration_status attribute under the aws_s3_bucket resource is now deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

  # The logging block_type under the aws_s3_bucket resource is now deprecated: Use the aws_s3_bucket_logging resource instead.
  logging = {}




  # The required parameter under the aws_s3_bucket resource was removed: .
  # The object_lock_enabled attribute under the aws_s3_bucket resource is now optional.
  # The object_lock_enabled attribute under the aws_s3_bucket resource is now deprecated: Use the top-level parameter object_lock_enabled instead.
  # An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
  object_lock_configuration {}
#  object_lock_enabled = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket" "c" {
  bucket = "my-tf-test-bucket"


  # The acceleration_status attribute under the aws_s3_bucket resource is now deprecated: Use the aws_s3_bucket_accelerate_configuration resource instead.
  acceleration_status = {}

  # The logging block_type under the aws_s3_bucket resource is now deprecated: Use the aws_s3_bucket_logging resource instead.
  logging = {}




  # The required parameter under the aws_s3_bucket resource was removed: .
  # The object_lock_enabled attribute under the aws_s3_bucket resource is now optional.
  # The object_lock_enabled attribute under the aws_s3_bucket resource is now deprecated: Use the top-level parameter object_lock_enabled instead.
  # An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
  object_lock_enabled = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.b.id
  acl    = "private"
}
data "aws_s3_account_public_access_block" "example" {
}

resource "aws_s3_bucket_website_configuration" "test" {
  bucket = aws_s3_bucket.b.bucket
}


