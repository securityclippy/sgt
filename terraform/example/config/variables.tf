variable "aws_profile" {
}
variable "aws_region" {}

variable "region" {
  default = "us-east-1"
}

variable "sgt_config_bucket_name" {}
variable "full_ssl_certchain" {}
variable "ssl_private_key" {}

variable "terraform_backend_bucket_name" {}

variable "environment" {}
variable "auto_approve_nodes" {}
variable "listen_address" {}
variable "domain" {}
variable "subdomain" {}
variable "email" {}
variable "renewal_threshold" {}
variable "use_le_staging" {}