output "pizza_auth_instance_ip" {
  description = "Public IP address of the pizza-auth microservice instance"
  value       = aws_instance.pizza_auth_instance.public_ip
}


