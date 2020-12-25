provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "terraform-santiago-terratest" {
  bucket = "terraform-santiago-terratest"
  acl    = "private"

  tags = {
    Name        = "terraform-santiago-terratest"
    Environment = "Dev"
  }
}

output "test_bucket" {
  value = aws_s3_bucket.terraform-santiago-terratest.bucket
}

resource "aws_s3_bucket_object" "object_test_1" {
  bucket  = "terraform-santiago-terratest"
  key     = "test1.txt"
  content = timestamp()
  depends_on = [
    aws_s3_bucket.terraform-santiago-terratest,
  ]
}

output "test1" {
  value = aws_s3_bucket_object.object_test_1.key

}

resource "aws_s3_bucket_object" "object_test_2" {
  bucket  = "terraform-santiago-terratest"
  key     = "test2.txt"
  content = timestamp()
  depends_on = [
    aws_s3_bucket.terraform-santiago-terratest,
  ]
}

output "test2" {
  value = aws_s3_bucket_object.object_test_2.key
}
