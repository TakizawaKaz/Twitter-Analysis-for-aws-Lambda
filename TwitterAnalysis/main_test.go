package main

import (
	"TwitterAnalysis/twitter"
	"context"
	"encoding/json"
	"fmt"
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

		var AssertionData ResponseJson
		err = json.Unmarshal(resp, &AssertionData)

		s := *AssertionData.Sentimental
		for i, v := range *AssertionData.Tweet {
			fmt.Printf(" Id:%s \n ScreenName:%s \n Name:%s \n CreatedAt: %v \n Text: %s ・・・ \n",
				v.Id, v.ScreenName, v.Name, v.CreatedAt, v.Text[0:50])
			fmt.Printf(" Sentiment:%s \n Positive:%.2f \n Negative:%.2f \n Mixed:%.2f \n Neutral:%.2f \n",
				s[i].Sentiment, s[i].Positive, s[i].Negative, s[i].Mixed, s[i].Neutral)
			fmt.Println()
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
