# awsgpt
Talk with AWS infratrucutre using natural laguage
## Project Description

This project is a Golang CLI application that leverages OpenAI's ChatGPT and AWS CLI to enable natural language chat with an AWS account and infrastructure.

## Prerequisites

Before running this application, make sure you have the following prerequisites installed:

- Golang: [Installation Guide](https://golang.org/doc/install)
- AWS CLI: [Installation Guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)


## This project

```
awsgpt/
├── cmd/
│   ├── awsgpt/
│   │   └── main.go
├── pkg/
│   ├── awsgpt/
│   │   ├── awsgpt.go
│   │   └── awsgpt_test.go
├── api/
│   ├── awsgpt/
│   │   ├── v1/
│   │   │   ├── awsgpt.pb.go
│   │   │   └── awsgpt.proto
├── web/
│   ├── static/
│   ├── templates/
├── scripts/
├── .gitignore
├── Dockerfile
├── Makefile
└── README.md

```

## Configuration

To use this application, you need to configure your AWS credentials. Follow the AWS CLI installation guide mentioned above to set up your credentials.

## Usage

To start a natural language chat session with your AWS account and infrastructure, run the following command:

