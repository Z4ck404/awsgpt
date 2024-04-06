package openai

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func NewClient(token string) *openai.Client {
	return openai.NewClient(token)
}

// GetAIResponse is a function that gets the response from the GPT-3.5 model
func GetAIResponse(cfg Config, masterPrompt string, messages []string) (string, error) {
	aiResponse := ""
	req := openai.ChatCompletionRequest{
		Model: cfg.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: masterPrompt,
			},
		},
	}

	for _, message := range messages {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		})
	}

	gptResponse, err := cfg.Client.CreateChatCompletion(context.Background(), req)
	aiResponse = gptResponse.Choices[0].Message.Content
	return aiResponse, err
}
