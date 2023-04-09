resource "aws_s3_bucket" "test" {
  bucket = var.name
}

resource "aws_s3_bucket_logging" "testlog" {
  bucket        = aws_s3_bucket.test.bucket
  target_bucket = ""
  target_prefix = ""
}

module "module_example_complete" {
  source  = "cloudposse/module/example//examples/complete"
  version = "1.0.0"
  # insert the 15 required variables here
}