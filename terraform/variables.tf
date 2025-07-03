variable "input"{
    default="kp"
    type=string
}

variable "name"{
    default="02"
    type=string
    description="name of the user"
}

output "filename" {
    value = var.name  
}

variable "names" {

  default = ["kp","pk","kk"]
}

output "names" {
    value = [for name in var.names: upper(name)]
}