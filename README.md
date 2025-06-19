# EventBridge to Redis Publisher

[![Go Version](https://img.shields.io/github/go-mod/go-version/denysandriyanov/eventbridge-to-redis-publisher)](https://golang.org/)

An AWS Lambda function that receives events from AWS EventBridge and publishes them to a Redis pub/sub channel.

## Overview

This Lambda function serves as a bridge between AWS EventBridge and Redis pub/sub messaging. It accepts JSON events from EventBridge triggers and forwards them to a specified Redis channel.

## Features

- Handles EventBridge events in AWS Lambda
- Publishes events to Redis pub/sub channels
- Supports TLS connections to Redis
- Configurable via environment variables

## Requirements

- Go 1.x
- AWS account with Lambda access
- Redis instance (accessible from Lambda)

## Configuration

The function requires the following environment variables:

| Variable | Description | Required |
|----------|-------------|----------|
| REDIS_HOST | Redis server address (host:port) | Yes |
| REDIS_TOPIC | Redis pub/sub channel name | Yes |
| REDIS_PASSWORD | Redis password | No |

## Building and Deploying

```powershell
# 1. Install the build-lambda-zip tool (optional if using PowerShell zip)
go install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest

# 2. Set environment variables for Linux build (required by Lambda)
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"

# 3. Build the binary with optimizations and name it 'bootstrap'
go build -ldflags="-s -w" -o bootstrap main.go

# 4. Package it into a Lambda-compatible zip file
Compress-Archive -Path bootstrap -DestinationPath function.zip

# 5. Deploy the zip file using your preferred method
# (e.g., AWS Console, Terraform, Serverless Framework, GitHub Actions, etc.)
```

## Usage

1. Configure an EventBridge rule to target this Lambda function
2. Set up your Redis consumers to subscribe to the configured topic
3. Events will flow from EventBridge → Lambda → Redis topic
