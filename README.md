# Terratest Example

This repo contains a Terratest example where Terraform provisions a small Ubuntu 16.04 instance and checks to verify that the outputted instance id matches with the instance id asserted through Terratest. For the sake of simplicity, we are using AWS as the targeted cloud provider.

## Terraform.tfvars file

The variable for creating the resource in AWS can be found in the `terraform.tfvars` file. The variable definitions are located in `variables.tf`.

### Variables.tf and Terraform.tfvars files

The definition of the variables are as follows:

```terraform
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
```

These defined variable types are used in the `terraform.tfvars` file as follows:
```terraform
# Specify path to AWS credentials
creds = "/home/hugh/.aws/credentials"

# Define the AWS region to deploy the test instance to
region = "us-west-2"


# Define the AMI and size of the instance
ami = "ami-08d70e59c07c61a3a"
instance = "t3.micro"
```
### NOTE
It is recommended that you change `creds` in `terraform.tfvars` to the location of your AWS credentials (`AWS_ACCESS_KEY` and `AWS_SECRET_KEY`)