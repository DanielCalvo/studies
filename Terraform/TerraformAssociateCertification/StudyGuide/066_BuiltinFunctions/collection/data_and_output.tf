//it looks like using echo to test variables isn't the best approach
//How about an output and local?
locals {
  myvar = "banana"
}
output "myoutput" {
//  value = local.myvar
  value = var.ami_id
}