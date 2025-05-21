terraform {
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "~> 2.0"
    }
  }

  backend "local" {
    path = "my-terr.tfstate"
  }
}


resource "local_file" "my_first_file" {
  content  = var.file_content
  filename = var.file_name
}


