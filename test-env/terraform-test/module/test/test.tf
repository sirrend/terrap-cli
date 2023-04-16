resource "aws_s3_bucket" "d" {
  bucket = "my-tf-test-bucket"
  alias = "test"


# deprecated: Use the aws_s3_bucket_logging resource instead.
  logging {}
  object_lock_configuration {

# A new 'object_lock_enabled' attribute is now available

# deprecated: Use the top-level parameter object_lock_enabled instead.

# The 'object_lock_enabled' attribute is now optional.

# The 'object_lock_enabled' attribute is no longer required.
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

# A new 'token' attribute is now available
  token = ""
  bucket = ""
  role = ""

# The 'nesting_mode' parameter has been changed from set to list
  rule {}
}
