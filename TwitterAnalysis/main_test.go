package main

import (
	"TwitterAnalysis/twitter"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	setup()

	code := m.Run()

	//teradown()
	os.Exit(code)
}
func setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func TestHandleRequest(t *testing.T) {

	t.Run("HandleRequest Events", func(t *testing.T) {

		tsAPIGwEvents := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"twquery": "#golang"},
		}
		ctx := context.Background()
		lc := &lambdacontext.LambdaContext{
			AwsRequestID: "testting request",
		}

		lambdacontext.NewContext(ctx, lc)
		ctx = lambdacontext.NewContext(ctx, lc)

		resp, err := HandleRequest(ctx, tsAPIGwEvents)

		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != 200{
			t.Errorf("Status Code != 200 : result = %v \n",resp.StatusCode)
		}

	})
}

func TestSearchTweet(t *testing.T) {
	config := oauth1.NewConfig(e.TwitterConsumerKey, e.TwitterConsumerSecret)
	token := oauth1.NewToken(e.TwitterAccessToken, e.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	_, err := twitter.SearchTweet(httpClient, "#golang")
	if err != nil {
		t.Fail()
	}
	//fmt.Println(string(result))
}
