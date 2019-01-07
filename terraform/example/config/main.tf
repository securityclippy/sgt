provider "aws" {
  profile = "${var.aws_profile}"
  region = "${var.aws_region}"
  version = ">= 1.21.0"
}

module "config" {
  source = "../../modules/config"
  aws_profile = "${var.aws_profile}"
  full_cert_chain = "${var.full_ssl_certchain}"
  priv_key = "${var.ssl_private_key}"
  terraform_backend_bucket_name = "${var.terraform_backend_bucket_name}"
  environment = "${var.environment}"
  auto_approve_nodes = "${var.auto_approve_nodes}"
  listen_address = "${var.listen_address}"
  domain = "${var.domain}"
  subdomain = "${var.subdomain}"
  email = "${var.email}"
  renewal_threshold = "${var.renewal_threshold}"
  use_le_staging = "${var.use_le_staging}"
}