terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure provider region
provider "aws" {
  region = "us-east-1"
}

# Create notification queue
resource "aws_sqs_queue" "notification_queue" {
  name = "notification-queue"
}

# Create notification topic
resource "aws_sns_topic" "user_updates" {
  name = "notification-topic"
}

