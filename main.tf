# Configure the s3 backend
terraform {
  backend "s3" {
    bucket = "example-bucket"
    key    = "example-key"
    region = "us-east-1"
  }
}

# Create an SQS queue
resource "aws_sqs_queue" "example" {
  name                      = "example-queue"
  delay_seconds             = 90
  max_message_size          = 256
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
}

# Create an API gateway
resource "aws_api_gateway_rest_api" "example" {
  name        = "example-api"
  description = "This is an example API"
}

# Create a DynamoDB table
resource "aws_dynamodb_table" "example" {
  name           = "example-table"
  billing_mode   = "PROVISIONED"
  read_capacity  = 10
  write_capacity = 10
  hash_key       = "id"

  attribute {
    name = "id"
    type = "S"
  }

  range_key      = "name"

}

resource "aws_lambda" "example" {
  function_name    = "example-function"
  handler          = "index.handler"
  runtime          = "nodejs12.x"
  role             = aws_iam_role.example.arn

}

provider "aws" {
  access_key = "AKIAIOSFODNN7EXAMPLE"
  secret_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
}
