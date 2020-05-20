package main

import (
	"TwitterAnalysis/analysis"
	"TwitterAnalysis/twitter"
	"TwitterAnalysis/utils"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/dghubble/oauth1"
	"strings"
)

var (
	e utils.Environment
)

type ResponseJson struct {
	Tweet       *[]twitter.TweetResult      `json:"tweet"`
	Sentimental *[]analysis.SentimentalData `json:"sentimental"`
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) ([]byte, error) {

	//環境変数を取得
	err := e.GetAccessKeys()
	if err != nil {
		return nil, err
	}

	//ツイッター検索ワードが設定されているかを判定
	q := event.QueryStringParameters
	var twQuery string
	if val, ok := q["twquery"]; ok {
		twQuery = val
	} else {
		return nil, utils.ErrorRaise("Error Non Query parameter [twquery]")
	}

	//TwitterAPI oauth1 での認証
	config := oauth1.NewConfig(e.TwitterConsumerKey, e.TwitterConsumerSecret)
	token := oauth1.NewToken(e.TwitterAccessToken, e.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	//ツイートサーチ
	twData, err := twitter.SearchTweet(httpClient, twQuery)
	if err != nil {
		return nil, err
	}

	//AWSセッション、認証情報をセット
	sess := session.Must(session.NewSession())
	cred := credentials.NewStaticCredentials(
		e.AwsAccessKeyId, e.AwsSecretAccessKey, e.AwsSessionToken)

	// Amazon Comprehend client の設定
	svc := comprehend.New(
		sess,
		aws.NewConfig().WithRegion(e.AwsRegion).WithCredentials(cred),
	)

	//センチメンタル分析用のクライアント設定
	textOnly := func() (text []string) {
		for _, tw := range twData {
			text = append(text, strings.Replace(tw.Text, "\n", "", -1))
		}
		return text
	}()
	client := analysis.NewAnalysis(svc, textOnly)
	//センチメンタル分析
	analysisData, err := client.SentimentalAnalysis()
	if err != nil {
		return nil, err
	}

	//レスポンス用データの定義
	respJson := ResponseJson{
		Tweet:       &twData,
		Sentimental: &analysisData,
	}

	//jsonで返答
	return json.Marshal(respJson)
}

func main() {
	lambda.Start(HandleRequest)
}
