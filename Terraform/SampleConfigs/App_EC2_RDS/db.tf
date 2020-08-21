resource "aws_db_instance" "dani" {
  allocated_storage    = 8
  storage_type         = "gp2"
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t2.micro"
  name                 = "mydb"
  username             = "admin"
  password             = "adminadmin"
  parameter_group_name = "default.mysql5.7"
  skip_final_snapshot  = "true"
}