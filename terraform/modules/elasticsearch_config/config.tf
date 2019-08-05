
data "terraform_remote_state" "firehose" {
  backend = "s3"
  config {
    bucket = "${var.terraform_backend_bucket_name}"
    key = "${var.environment}/elasticsearch_firehose/terraform.tfstate"
    profile = "${var.aws_profile}"
    region = "${var.aws_region}"
  }
}

data "terraform_remote_state" "datastore" {
  backend = "s3"
  config {
    bucket = "${var.terraform_backend_bucket_name}"
    key = "${var.environment}/datastore/terraform.tfstate"
    profile = "${var.aws_profile}"
    region = "${var.aws_region}"
  }
}

data "aws_route53_zone" "osquery-sgt-dns-zone" {
  name = "${var.domain}"
}

resource "aws_s3_bucket_object" "osquery-sgt-binary" {
  bucket = "${aws_s3_bucket.sgt_config_s3_bucket.bucket}"
  source = "../../../sgt"
  key = "sgt/sgt"
  etag = "${md5(file("../../../sgt"))}"
}

data "template_file" "sgt-config-file" {
  template = "${file("${path.module}/example.config.json")}"
  vars {
    firehose_aws_access_key_id = "${data.terraform_remote_state.firehose.sgt-node-user-access-key-id}"
    firehose_aws_secret_access_key = "${data.terraform_remote_state.firehose.sgt-node-user-secret-access-key}",
    firehose_stream_name = "${data.terraform_remote_state.firehose.sgt-firehose-stream-name}",
    distributed_query_logger_firehose_stream_name = "${data.terraform_remote_state.firehose.sgt-distributed-firehose-stream-name}"
    auto_approve_nodes = "${var.auto_approve_nodes}"
    domain = "${var.subdomain}.${var.domain}"
    email = "${var.email}"
    hosted_zone_id = "${data.aws_route53_zone.osquery-sgt-dns-zone.zone_id}"
    #s3_config_bucket = "${data.terraform_remote_state.datastore.s3_bucket_name}"
    s3_config_bucket = "${aws_s3_bucket.sgt_config_s3_bucket.bucket}"
    listen_address = "${var.listen_address}"
    use_le_staging = "${var.use_le_staging}"
  }
}

resource "aws_kms_key" "sgt_config_bucket_kms_key" {
  deletion_window_in_days = 10
}

resource "aws_s3_bucket" "sgt_config_s3_bucket" {
  bucket = "${var.sgt_config_bucket}"
}

resource "aws_s3_bucket_object" "osquery-sgt-config" {
  bucket = "${aws_s3_bucket.sgt_config_s3_bucket.bucket}"
  content = "${data.template_file.sgt-config-file.rendered}"
  key = "sgt/config.json"
  etag = "${md5("{data.template_file.sgt-config-file.rendered}")}"
}

