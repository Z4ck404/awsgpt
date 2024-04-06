package aws

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	DEFAULT_AWS_PROFILE = "default"
	DEFAULT_AWS_REGION  = "us-east-1"
)

type Config struct {
	Profile string
	Region  string
}

// RunCommand takes an aws command and runs it then returns the output
func RunCommand(cfg Config, command string) (string, error) {
	formatedCommand := formatAWSCommand(command)
	log.Println("Running command: ", formatedCommand)
	cmd := exec.Command("sh", "-c", formatedCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command: %w", err)
	}
	return string(output), nil
}

func formatAWSCommand(command string) string {
	return fmt.Sprintf(command)
}
