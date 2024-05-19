package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/DavidVergison/simplesheet/bot/discordlib"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	rawBody := request.Body
	fmt.Println("start")
	fmt.Println(rawBody)

	// validate nacl header
	err := validator.VerifyRequest(
		request.Headers["x-signature-ed25519"],
		request.Headers["x-signature-timestamp"],
		request.Body,
	)
	if err != nil {
		fmt.Println("StatusUnauthorized")
		return formatUnauthorizedResponse()
	}

	discordRequest := discordlib.DiscordRequestDto{}
	err = json.Unmarshal([]byte(rawBody), &discordRequest)

	// validate ping
	fmt.Println("test ping")
	if validator.IsPing(discordRequest.Type) {
		fmt.Println("is ping !")
		return formatPingResponse()
	}

	return discordSlashCommandHandler(discordRequest)
}

func formatPingResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       discordlib.GetPingResponse(),
		StatusCode: 200,
	}, nil
}

func formatResponse(msg string) (events.APIGatewayProxyResponse, error) {
	rawResponse := discordlib.DiscordResponseDto{
		Type: 4,
		Data: discordlib.DiscordDataResponseDto{
			Content: msg,
		},
	}

	jsonResponse, _ := json.Marshal(rawResponse)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResponse),
		StatusCode: 200,
	}, nil
}

func formatUnauthorizedResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusUnauthorized,
	}, nil
}

var validator discordlib.DiscordCommandValidator
func main() {
	validator = discordlib.DiscordCommandValidator{
		PublicKey: "b33dc913471e37e35cf194f5e0b8efddac3b208149fd231154a6aa537a93376e",
	}
	lambda.Start(handler)
}
