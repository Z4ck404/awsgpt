# [WIP] awsgpt

<img width="1089" alt="Capture d’écran 2024-04-07 à 06 15 15" src="https://github.com/Z4ck404/awsgpt/assets/35115877/40322d5b-bb64-46fa-933d-9bed1e5d9866">


## What

This project is a Golang CLI application that leverages OpenAI's ChatGPT and AWS CLI to enable natural language chat with an AWS account and infrastructure.

## Prerequisites

Before running this application, make sure you have the following prerequisites installed:

- Golang: [Installation Guide](https://golang.org/doc/install)
- AWS CLI: [Installation Guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

## Usage

1. Get your OpenAI token from the [OpenAI website](https://platform.openai.com/account/api-keys).
2. [OPTIONAL] Set the `OPENAI_API_KEY` environment variable:

    ```bash
    export OPENAI_API_KEY="sk-<OPEN_AI_TOKEN>"
    ```

3. Build the CLI:

    ```bash
    make build
    ```

4. Chat with your AWS Account:

    - The CLI will use the default profile to connect to your AWS account.

    ```bash
    ./bin/awsgpt --token="sk-<OPEN_AI_TOKEN>" --question "Do I have any buckets in my account?"
    ```

    ```bash
    ./bin/awsgpt -h
    ```
