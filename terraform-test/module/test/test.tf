resource "aws_s3_bucket" "d" {
  bucket = "my-tf-test-bucket"
  alias = "test"


# deprecated: Use the aws_s3_bucket_logging resource instead.
  logging {}

  object_lock_configuration {

# An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.

# deprecated: Use the top-level parameter object_lock_enabled instead.

# The object_lock_enabled attribute under the aws_s3_bucket resource is now optional.

# The required parameter under the aws_s3_bucket resource was removed and can no longer be used.
    object_lock_enabled = ""
  }

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}


# A new resource named aws_s3_bucket_acl is now available
resource "aws_s3_bucket_acl" "example13" {
  bucket = aws_s3_bucket.d.id
  acl    = "private"
}

data "aws_s3_account_public_access_block" "example2d" {
}

resource "aws_s3_bucket_replication_configuration" "test" {

# An 'token' attribute was added under the aws_s3_bucket_replication_configuration resource.
  token = ""
  bucket = ""
  role = ""

# The type of the nesting_mode parameter component was changed from set to list
  rule {}
}
