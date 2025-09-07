package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"github.com/vietnguyen-dev/go-server/routes/models"
)

func GenerateReport(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		http.Error(w, "no api key for openai is not set", http.StatusInternalServerError)
		return
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)

	var reportRequest models.ReportRequest
	err := json.NewDecoder(r.Body).Decode(&reportRequest)
	if err := reportRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonText, err := json.Marshal(reportRequest.MoodData)
	if err != nil {
		panic(err)
	}
	reportText := fmt.Sprintf(`Using this data %s,
	create me a report of how my moods have been

	include: 
	- highest, lowest, mean, median, mode mood
	- a brief summary of how its been going
	- what are some of my triggers
	- what are some things that make my mood better
	`,
		string(jsonText))

	stream := client.Chat.Completions.NewStreaming(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(reportText),
		},
		Seed:  openai.Int(0),
		Model: openai.ChatModelGPT5Nano,
	})

	// optionally, an accumulator helper can be used
	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		if content, ok := acc.JustFinishedContent(); ok {
			println("Content stream finished:", content)
		}

		// if using tool calls
		if tool, ok := acc.JustFinishedToolCall(); ok {
			println("Tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
		}

		if refusal, ok := acc.JustFinishedRefusal(); ok {
			println("Refusal stream finished:", refusal)
		}

		// it's best to use chunks after handling JustFinished events
		if len(chunk.Choices) > 0 {
			println(chunk.Choices[0].Delta.Content)
		}
	}

	if stream.Err() != nil {
		panic(stream.Err())
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(acc.Choices[0].Message.Content))
}
