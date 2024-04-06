package cmd

import (
	"log"
	"os"

	"awsgpt/internal/aws"
	openaiclient "awsgpt/internal/openai"

	"github.com/spf13/cobra"

	"github.com/sashabaranov/go-openai"
)

// openaiConfig is the configuration for the openai client
var openaiConfig = &openaiclient.Config{
	Model: openai.GPT3Dot5Turbo,
}

// awsConfig is the configuration for the aws CLI
var awsConfig = &aws.Config{
	Profile: aws.DEFAULT_AWS_PROFILE,
	Region:  aws.DEFAULT_AWS_REGION,
}

var rootCmd = &cobra.Command{
	Use:   "awsgpt",
	Short: "Talk with your AWS account using gpt",
	Long: `
		awsgpt is a go CLI library that allows you to 
		talk with your AWS account using gpt-3.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Execute()
	},
}

func init() {
	// Initialize the config
	cobra.OnInitialize(initConfig)

	// Add the flags to the root command
	rootCmd.PersistentFlags().StringVar(&awsConfig.Profile, "aws-profile", "", "your aws profile (default is Default)")
	rootCmd.PersistentFlags().StringVar(&awsConfig.Profile, "aws-region", "", "the default aws region for aws CLI if not set in the aws profile (default is us-east-1)")
	rootCmd.PersistentFlags().StringVar(&openaiConfig.Token, "token", "", "the openai token to use (default is the environment variable OPENAI_API_TOKEN)")
	// Add the message flag to the root command
	rootCmd.PersistentFlags().StringP("message", "a", "YOUR NAME", "Author name for copyright attribution")

	//TODO Enable verbose logging
	//rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose logging")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	// Get the messages from the flags
	userQuestion := rootCmd.Flag("message").Value.String()
	userQuestionMessages := []string{userQuestion}

	// Get the AI response
	aiResponse, err := openaiclient.GetAIResponse(*openaiConfig, openaiclient.COMMAND_MASTER_PROMPT, userQuestionMessages)
	if err != nil {
		log.Fatalf("Failed to get AI response: %v", err)
		os.Exit(1)
	}

	log.Printf("AI Response: %s\n", aiResponse)
	// Run the command
	awsCommand := aiResponse
	awsOutput, err := aws.RunCommand(*awsConfig, awsCommand)
	if err != nil {
		log.Fatalf("Failed to run the AWS command: %v", err)
		os.Exit(1)
	}

	// Print the output
	log.Printf("AWS Command: %s\n", awsCommand)
	log.Printf("AWS Output: %s\n", awsOutput)

	// Explain the output with AI
	awsCommandOutputMessage := []string{awsOutput}
	aiExplainCommand, err := openaiclient.GetAIResponse(*openaiConfig, openaiclient.SUMMARIZE_CLI_OUTPUT_COMMAND_PROMPT, awsCommandOutputMessage)
	if err != nil {
		log.Fatalf("Failed to get AI response: %v", err)
		os.Exit(1)
	}

	log.Printf("%s\n", aiExplainCommand)
}

func initConfig() {

	// Get the openai from the --token flag if set and if not get it from the environment variable
	token := rootCmd.Flag("token").Value.String()
	if token != "" {
		openaiConfig.Token = token
	} else {
		openAItoken, err := openaiclient.GetOpenAITokenEnvVar()
		if err != nil {
			log.Fatalf("Failed to get OpenAI token: %v", err)
			os.Exit(0)
		}
		openaiConfig.Token = openAItoken
	}

	// Initialize the openai client
	openaiConfig.Client = openaiclient.NewClient(openaiConfig.Token)

	// Initialize the aws config
	awsConfig.Profile = rootCmd.Flag("aws-profile").Value.String()
	awsConfig.Region = rootCmd.Flag("aws-region").Value.String()

}
