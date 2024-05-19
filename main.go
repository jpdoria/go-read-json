package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Records struct which contains an array of records.
type Records struct {
	Records []Record `json:"Records"`
}

// Record struct which contains event data.
type Record struct {
	Sns Sns `json:"Sns"`
}

// Sns struct which contains SNS data like Message, MessageId, Type, etc.
type Sns struct {
	Message string `json:"Message"`
}

// Message struct which contains the message from the trigger. In this case, it's a message from CodePipeline.
type Message struct {
	Region string `json:"region"`
	Detail Detail `json:"detail"`
}

// Detail struct which contains the CodePipeline state change details.
type Detail struct {
	Pipeline    string `json:"pipeline"`
	ExecutionId string `json:"execution-id"`
	State       string `json:"state"`
}

func main() {
	// Open the JSON file and get the byte array.
	fileName := "event.json"
	byteValue, err := os.ReadFile(fileName)

	// Print the error if the file couldn't be opened.
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message if the file was opened successfully, and defer the closing of the file.
	log.Printf("%q file opened successfully!\n", fileName)

	// Initialize the Records struct.
	var records Records

	// Unmarshal the byte array into the Records struct.
	json.Unmarshal(byteValue, &records)

	// Initialize the Message struct.
	var message Message

	// Unmarshal the Message field from the Sns struct into the Message struct.
	msg, err := strconv.Unquote("`" + records.Records[0].Sns.Message + "`")

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(msg), &message)

	// Print the message from the trigger (CodePipeline).
	log.Printf("Region: %v", message.Region)
	log.Printf("Pipeline Name: %v", message.Detail.Pipeline)
	log.Printf("Execution ID: %v", message.Detail.ExecutionId)
	log.Printf("State: %v", message.Detail.State)

	// Print the complete URL to the CodePipeline execution and the message for Slack or Google Chat.
	url := fmt.Sprintf("https://%v.console.aws.amazon.com/codesuite/codepipeline/pipelines/%v/executions/%v/visualization?region=%v", message.Region, message.Detail.Pipeline, message.Detail.ExecutionId, message.Region)
	log.Printf("URL: %v", url)
	chatMessage := fmt.Sprintf("%v has %v. Please check this link for more details: %v", message.Detail.Pipeline, strings.ToLower(message.Detail.State), url)
	log.Printf("Message for Slack or Google Chat: %v", chatMessage)
}
