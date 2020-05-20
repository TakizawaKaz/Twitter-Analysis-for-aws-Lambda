package main

import (
	"TwitterAnalysis/analysis"
	"TwitterAnalysis/twitter"
	"TwitterAnalysis/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response:= events.APIGatewayProxyResponse{}

	//環境変数を取得
	err := e.GetAccessKeys()
	if err != nil {
		return response, err
	}

	//ツイッター検索ワードが設定されているかを判定
	q := event.QueryStringParameters
	var twQuery string
	if val, ok := q["twquery"]; ok {
		twQuery = val
	} else {
		return response, utils.ErrorRaise("Error Non Query parameter [twquery]")
	}

	fmt.Printf("query word %s \n",q)

	//TwitterAPI oauth1 での認証
	config := oauth1.NewConfig(e.TwitterConsumerKey, e.TwitterConsumerSecret)
	token := oauth1.NewToken(e.TwitterAccessToken, e.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	//ツイートサーチ
	twData, err := twitter.SearchTweet(httpClient, twQuery)
	if err != nil {
		return response, err
	}

	fmt.Printf("Search tweet Seccess : %v \n",len(twData))

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
		return response, err
	}

	fmt.Printf("Amazon Comprehend SentimentalAnalysis Seccess : %v \n",len(analysisData))

	//レスポンス用データの定義
	respJson := ResponseJson{
		Tweet:       &twData,
		Sentimental: &analysisData,
	}

	//Json型に変換
	resp,err:=json.Marshal(respJson)
	if err != nil {
		return response, err
	}

	//ログ用にインデントを入れたJsonを作成
	var indentJson bytes.Buffer
	err = json.Indent(&indentJson, resp, "", "  ")
	if err != nil {
		return response, err
	}
	fmt.Printf("retrun data :\n %v \n",string(indentJson.Bytes()))

	//jsonで返答
	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	},nil
}

func main() {
	lambda.Start(HandleRequest)
}
