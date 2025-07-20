resource "aws_db_instance" "main" {
  allocated_storage    = 20
  engine               = "postgres"
  engine_version       = "13.7"
  instance_class       = "db.t3.micro"
  name                 = "autonomouspen"
  username             = "postgres"
  password             = var.db_password
  parameter_group_name = "default.postgres13"
  skip_final_snapshot  = true
  vpc_security_group_ids = [aws_security_group.rds.id]
  db_subnet_group_name = aws_db_subnet_group.main.name
}

resource "aws_db_subnet_group" "main" {
  name       = "main"
  subnet_ids = var.subnet_ids
}

resource "aws_security_group" "rds" {
  name        = "rds-sg"
  description = "Allow traffic to RDS"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Replace with more restrictive CIDR blocks
  }
}

variable "db_password" {
  type      = string
  sensitive = true
}

variable "subnet_ids" {
  type = list(string)
}

variable "vpc_id" {
  type = string
}
