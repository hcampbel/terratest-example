terraform {
  # This module is now only being tested with Terraform 0.13.x. However, to make upgrading easier, we are setting
  # 0.12.26 as the minimum version, as that version added support for required_providers with source URLs, making it
  # forwards compatible with 0.13.x code.
  required_version = ">= 0.12.26"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }

}

provider "aws" {
  region                  =  var.region
  shared_credentials_file = var.creds
}

resource "aws_instance" "example" {
  ami           = var.ami
  instance_type = var.instance

  tags = {
      Name = "Test Instance"
  }

}

output "id" {
    value = aws_instance.example.id
}

