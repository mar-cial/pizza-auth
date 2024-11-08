# main.tf

resource "aws_security_group" "pizza_auth_sg" {
  name_prefix = "pizza-auth-sg-"
  description = "Security group for the pizza-auth microservice"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "pizza_auth_instance" {
  ami           = var.ami_id
  instance_type = var.instance_type
  key_name      = var.ssh_key_name

  security_groups = [aws_security_group.pizza_auth_sg.name]

  tags = {
    Name = "pizza-auth-microservice"
  }
}


