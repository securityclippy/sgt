variable "aws_profile" {
}

variable "aws_region" {}

variable "aws_keypair" {}

variable "domain" {}

variable "subdomain" {}

variable "asg_min_size"  {
  default = 2
}

variable "asg_max_size" {
  default = 4
}

variable "asg_desired_size" {}

variable "terraform_backend_bucket_name" {}

variable "environment" {}
variable "sgt_config_bucket_name" {}