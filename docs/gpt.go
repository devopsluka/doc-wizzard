package docs

import (
	"log"
	"os"

	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"github.com/devopsluka/doc-wizzard/types"
	"github.com/joho/godotenv"
)

func CleanSheet(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	if err := file.Truncate(0); err != nil {
		log.Fatalf("Error truncating file: %v", err)
	}
}

func Append2File(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func CallGPT(obj types.K8sObject) (string, error) {
	var error = godotenv.Load(".env")
	if error != nil {
		log.Fatalf("Error loading .env file: %v", error)
	}

	var API_KEY = os.Getenv("OPENAI_API_KEY")
	if API_KEY == "" {
		log.Fatalf("OPENAI_API_KEY is not set: %v", error)
	}

	doc, err := GenerateK8sObjectDocumentation(obj)
	if err != nil {
		log.Fatalf("Error generating documentation: %v", err)
	}

	c, _ := gpt35.NewClient(API_KEY)
	req := &gpt35.Request{
		Model:       gpt35.ModelGpt35Turbo,
		Temperature: 0.2,
		MaxTokens:   500,
		Messages: []*gpt35.Message{
			{
				Role:    gpt35.RoleUser,
				Content: "Please write documentation for this Go template YAML manifest: " + doc,
			},
		},
	}

	resp, err := c.GetChat(req)
	if err != nil {
		return "", err
	}

	Append2File("output/documentation.txt", resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content, nil

}
