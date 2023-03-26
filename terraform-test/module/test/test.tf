resource "aws_s3_bucket" "d" {
  bucket = "my-tf-test-bucket"
  alias = "test"





# The required parameter under the aws_s3_bucket resource was removed: .




# The required parameter under the aws_s3_bucket resource was removed: .




# The required parameter under the aws_s3_bucket resource was removed: .
# The object_lock_enabled attribute under the aws_s3_bucket resource is now optional.
# The object_lock_enabled attribute under the aws_s3_bucket resource is now deprecated: Use the top-level parameter object_lock_enabled instead.
# An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
# The object_lock_enabled attribute under the aws_s3_bucket resource is now optional.
# The object_lock_enabled attribute under the aws_s3_bucket resource is now deprecated: Use the top-level parameter object_lock_enabled instead.
# An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
# The object_lock_enabled attribute under the aws_s3_bucket resource is now optional.
# The object_lock_enabled attribute under the aws_s3_bucket resource is now deprecated: Use the top-level parameter object_lock_enabled instead.
# An 'object_lock_enabled' attribute was added under the aws_s3_bucket resource.
  object_lock_enabled = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket_acl" "example13" {
  bucket = aws_s3_bucket.d.id
  acl    = "private"
}

data "aws_s3_account_public_access_block" "example2d" {
}
