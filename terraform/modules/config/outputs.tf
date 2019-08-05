output "sgt_config_bucket_arn" {
  value = "${aws_s3_bucket.sgt_config_s3_bucket.arn}"
}