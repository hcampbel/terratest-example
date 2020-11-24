variable "creds" {
    type = string   
}

variable "region" {
    type = string
    default = "us-east-1"
}

variable "ami" {
    type = string
}

variable "instance" {
    type = string
    default = "t2.micro"
}
