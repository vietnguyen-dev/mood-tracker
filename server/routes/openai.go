package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func GenerateReport(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Generating report")
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		http.Error(w, "OPENAI_API_KEY is not set", http.StatusInternalServerError)
		return
	}
	question := r.URL.Query().Get("question")
	if question == "" {
		http.Error(w, "question is required", http.StatusBadRequest)
		return
	}
	client := openai.NewClient(
		option.WithAPIKey(apiKey), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)
	response, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT5Nano,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		},
	})
	if err != nil {
		panic(err.Error())
	}
	println(response.Choices[0].Message.Content)
}