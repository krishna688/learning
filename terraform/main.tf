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

provider "local" {
  
}

resource "local_file" "test" {
  filename = "test_${var.name}.txt"
  content = var.input
}





