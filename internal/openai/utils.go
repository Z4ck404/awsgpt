package openai

import (
	"fmt"
	"os"
)

func GetOpenAITokenEnvVar() (string, error) {
	return getenvVar(OPENAI_API_TOKEN_VAR_NAME)
}

func getenvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s environment variable not set", key)
	}
	return value, nil
}
