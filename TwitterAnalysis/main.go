package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
	"time"
)

type TweetResult struct {
	Id         string
	ScreenName string
	Name       string
	CreatedAt  time.Time
	text       string
	Positive   float64
	Negative   float64
	Mixid      float64
	Natural    float64
}

const (
	TwitterTimeLayout = "Mon Jan 2 15:04:05 -0700 2006"
)
var (
	e Environment
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func init() {
	err := e.GetAccessKeys()
	ErrorCheck(err)
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (response []byte, err error) {

	q := event.QueryStringParameters

	var twQuery string
	if val, ok := q["twquery"]; ok {
		twQuery = val
	} else {
		return nil, ErrorRaise("Error Non Query parameter [twquery]")
	}

	config := oauth1.NewConfig(e.TwitterConsumerKey, e.TwitterConsumerSecret)
	token := oauth1.NewToken(e.TwitterAccessToken, e.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	err = SearchTweet(httpClient, twQuery)
	ErrorCheck(err)

	return nil, err

}
func SearchTweet(httpClient *http.Client, query string) error {

	// Twitter client
	client := twitter.NewClient(httpClient)

	//検索パラメーターをセット
	twitterParam := twitter.SearchTweetParams{
		Query:      fmt.Sprintf("%s -filter:retweets", query),
		Lang:       "ja",
		Locale:     "ja",
		ResultType: "recent",
	}

	//APIの結果
	search, resp, err := client.Search.Tweets(&twitterParam)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return ErrorRaise(fmt.Sprintf("HTTP Statas code :%v \n", resp.Status))
	}

	twr := make([]TweetResult,0,len(search.Statuses))
	for _, tweet := range search.Statuses {

		crAt, _ := time.Parse(TwitterTimeLayout, tweet.CreatedAt)
		t := TweetResult{
			Id:         tweet.User.IDStr,
			ScreenName: tweet.User.ScreenName,
			Name:       tweet.User.Name,
			CreatedAt:  crAt.In(jst),
			text:       tweet.Text,
			Positive:   0,
			Negative:   0,
			Mixid:      0,
			Natural:    0,
		}

		fmt.Println(t)
		twr = append(twr, t)
	}

	return nil //Tweetオブジェクトを返す
}
