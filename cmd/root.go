package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"awsgpt/internal/aws"
	openaiclient "awsgpt/internal/openai"

	"github.com/spf13/cobra"
)

var verbose bool

// openaiConfig is the configuration for the openai client
var openaiConfig = &openaiclient.Config{}

// awsConfig is the configuration for the aws CLI
var awsConfig = &aws.Config{
	Profile: aws.DEFAULT_AWS_PROFILE,
	Region:  aws.DEFAULT_AWS_REGION,
}

var rootCmd = &cobra.Command{
	Use:   "awsgpt",
	Short: "[AWSGPT] Talk with your AWS account using gpt",
	Long: `
		awsgpt is a go CLI library that allows you to 
		talk with your AWS account using gpt-3.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Root().Short)
	},
}

func init() {
	// Initialize the config
	cobra.OnInitialize(initConfig)

	// Add the aws flags to the root command
	rootCmd.PersistentFlags().StringVar(&awsConfig.Profile, "aws-profile", "", "your aws profile (default is Default)")
	rootCmd.PersistentFlags().StringVar(&awsConfig.Region, "aws-region", "", "the default aws region for aws CLI if not set in the aws profile (default is us-east-1)")

	// Add the token and Model flags to the root command
	rootCmd.PersistentFlags().StringVar(&openaiConfig.Token, "token", "", "the openai token to use (default is the environment variable OPENAI_API_TOKEN)")
	//rootCmd.PersistentFlags().StringVar(&openaiConfig.Model, "model", "", "the openai model to use (default is gpt-3.5-turbo model)")

	// Add the input message flag to the root command
	rootCmd.PersistentFlags().StringP("question", "", "Your question", "The question you want to ask the AI model about your AWS account.")

	//Enable verbose logging
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	// Enable verbose logging if the verbose flag is set
	if verbose {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}

	// Get the messages from the flags
	userQuestion := rootCmd.Flag("question").Value.String()
	userQuestionMessages := []string{userQuestion}

	// Get the AI response
	aiResponse, err := openaiclient.GetAIResponse(*openaiConfig, openaiclient.COMMAND_MASTER_PROMPT, userQuestionMessages)
	if err != nil {
		log.Fatalf("Failed to get AI response: %v", err)
		os.Exit(1)
	}

	log.Printf("[AWS CLI] AI Response: %s\n", aiResponse)
	// Run the command
	awsCommand := aiResponse
	awsOutput, err := aws.RunCommand(*awsConfig, awsCommand)
	if err != nil {
		log.Fatalf("Failed to run the AWS command: %v", err)
		os.Exit(1)
	}

	// Print the output
	//log.Printf("AWS Command: %s\n", awsCommand)
	log.Printf("[AWS CLI OUTPUT] AWS Output: %s\n", awsOutput)

	// Explain the output with AI
	awsCommandOutputMessage := []string{awsOutput}
	aiExplainCommand, err := openaiclient.GetAIResponse(*openaiConfig, openaiclient.SUMMARIZE_CLI_OUTPUT_COMMAND_PROMPT, awsCommandOutputMessage)
	if err != nil {
		log.Fatalf("Failed to get AI response: %v", err)
		os.Exit(1)
	}

	log.Printf(" [AWS CLI EXPLAINED] %s\n", aiExplainCommand)

	fmt.Println(aiExplainCommand)
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
	if rootCmd.Flag("aws-profile").Value.String() == "" {
		awsConfig.Profile = aws.DEFAULT_AWS_PROFILE
	} else {
		awsConfig.Profile = rootCmd.Flag("aws-profile").Value.String()
	}

	if rootCmd.Flag("aws-region").Value.String() == "" {
		awsConfig.Region = aws.DEFAULT_AWS_REGION
	} else {
		awsConfig.Region = rootCmd.Flag("aws-region").Value.String()
	}

	// validate the user question
	userQuestion := rootCmd.Flag("question").Value.String()
	if userQuestion == "" {
		fmt.Errorf("Please provide a message to ask the AI mode: %s")
		os.Exit(1)
	}

}
