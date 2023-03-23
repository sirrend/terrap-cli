resource "aws_s3_bucket" "c" {
  bucket = "my-tf-test-bucket"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket_acl" "example1" {
  bucket = aws_s3_bucket.c.id
  acl    = "private"
}

data "aws_s3_account_public_access_block" "example2" {
}