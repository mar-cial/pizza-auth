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
