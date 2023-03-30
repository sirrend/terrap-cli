resource "aws_s3_bucket" "d" {
  bucket = "my-tf-test-bucket"
  alias = "test"

  logging {}
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
