resource "aws_key_pair" "mykey"{
  key_name = "mykey"
  public_key = ""

}

resource "aws_instance" "daniinstance"{
  provisioner "file" {
    source = "script.sh"
    destination = "/tmp/script.sh"
  }
  provisioner "remote-exec" {
    inline = [
    "chmod +x /tmp/script.sh",
     "sudo /tmp/script"
    ]
  }
  connection {
    user = "${var.INSTANCE_USERNAME}"
    private_key = "${file("${var.PATH_TO_PRIVATE_KEY}")}"
  }
}