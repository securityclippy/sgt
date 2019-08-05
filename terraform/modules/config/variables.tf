

variable "aws_profile" {}

variable "aws_region" {
  default = "us-east-1"
}

variable "terraform_backend_bucket_name" {}

variable "environment" {}
variable "auto_approve_nodes" {}

variable "domain" {}
variable "subdomain" {}
variable "email" {}
variable "listen_address" {}
variable "use_le_staging" {}
variable "sgt_config_bucket" {}
