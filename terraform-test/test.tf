resource "aws_s3_bucket" "b" {
  bucket = "my-tf-test-bucket"

  acceleration_status = {}
  logging = {}
  object_lock_enabled = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
resource "aws_s3_bucket" "fasdfasd" {
  bucket = "my-tf-test-bucket"

  acceleration_status = {}
  logging = {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket" "d" {
  bucket = "my-tf-test-bucket"

  acceleration_status = {}
  logging = {}

  object_lock_configuration {}

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket" "c" {
  bucket = "my-tf-test-bucket"

  acceleration_status = {}
  logging = {}
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


