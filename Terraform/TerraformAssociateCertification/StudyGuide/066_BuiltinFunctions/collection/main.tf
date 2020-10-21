//
//resource "null_resource" "just_nulling" {
//  provisioner "local-exec"{
//    //lookup(map, key, default)
//    command = "echo ${lookup(var.ami_id, "ubuntu", "banana")}"
//  }
//  provisioner "local-exec"{
//    command = "echo ${lookup(var.ami_id, "aaaaaaaaaaaaaa", "banana")}"
//  }
//}
//
//resource "null_resource" "just_nulling2" {
//  provisioner "local-exec"{
//    //coalesce: takes any number of arguments and returns the first one that isn't null or an empty string  (needs to take strings, not lists of strings)
//    command = "echo ${coalesce("", "a","b","c")}"
//  }
//}
//
////Couldn't make this work yet
////resource "null_resource" "just_nulling3" {
////  provisioner "local-exec"{
////    //coalesce: takes any number of arguments and returns the first one that isn't null or an empty string  (needs to take strings, not lists of strings)
////    command = "echo ${coalescelist(["a", "b"], ["c", "d"])}"
////  }
////}
//
////Couldn't make this work yet
//resource "null_resource" "just_nulling4" {
//  provisioner "local-exec"{
//    //coalesce: takes any number of arguments and returns the first one that isn't null or an empty string  (needs to take strings, not lists of strings)
//    command = "echo ${var.myotherlist[0]}" //Oh, this cannot echo a list
//  }
//}
//
//
//
