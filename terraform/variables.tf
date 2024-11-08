variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "us-west-1"
}

variable "instance_type" {
  description = "Instance type of EC2"
  type        = string
  default     = "t2.micro"
}

variable "ami_id" {
  description = "AMI ID"
  type        = string
}

variable "ssh_key_name" {
  description = "ssh key name"
  type        = string
}
