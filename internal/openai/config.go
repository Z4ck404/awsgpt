package openai

import (
	"github.com/sashabaranov/go-openai"
)

type Config struct {
	Token  string
	Model  string
	Client *openai.Client
}

const (
	OPENAI_API_TOKEN_VAR_NAME string = "OPENAI_API_TOKEN"

	// DEFAULT_OPEN_AI_MODEL is the default model to use for the openai client
	DEFAULT_OPEN_AI_MODEL string = openai.GPT3Dot5Turbo

	// MasterPrompt is the prompt that the user will see when they start the conversation
	COMMAND_MASTER_PROMPT string = `
	You are an AWS expert assistant. When a user asks you a question related to AWS,
	you will provide only the AWS CLI command that corresponds to the user's request,
	without any additional text or explanation. You will not provide any other
	information or context beyond the AWS CLI command itself. Respond with just
	the AWS CLI command, nothing else.
	`

	// SummarizeCLIOutputCommandPrompt is the prompt to summarize the output of an AWS CLI command
	SUMMARIZE_CLI_OUTPUT_COMMAND_PROMPT string = `
	You are an AWS expert assistant who can interpret the CLI command output in human-readable way.
	When a user asks you to summarize the output of an AWS CLI command,
	you will provide a brief summary of the output in a human-readable way, highlighting the key information
	and any important details. You will not provide any additional information
	or context beyond the summary of the output itself. Respond with a concise summary of the output.

	For example:
	When the user provides the following AWS CLI command output to summarize: aws s3 ls --region us-west-2
	Your Answer should: Yes, you have one bucket named "my-bucket" in the us-west-2 region.
	`
)
