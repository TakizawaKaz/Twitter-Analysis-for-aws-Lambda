package twitter

import (
	"TwitterAnalysis/utils"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"net/http"
	"strings"
	"time"
)

type TweetResult struct {
	Id         string    `json:"id"`
	ScreenName string    `json:"screen_name"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	Text       string    `json:"text"`
}

const (
	TimeLayout = "Mon Jan 2 15:04:05 -0700 2006"
)

var (
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func SearchTweet(httpClient *http.Client, query string) ([]TweetResult, error) {

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
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, utils.ErrorRaise(fmt.Sprintf("HTTP Statas code :%v \n", resp.Status))
	}

	twr := make([]TweetResult, 0, len(search.Statuses))
	for _, tweet := range search.Statuses {

		crAt, _ := time.Parse(TimeLayout, tweet.CreatedAt)
		t := TweetResult{
			Id:         tweet.User.IDStr,
			ScreenName: tweet.User.ScreenName,
			Name:       tweet.User.Name,
			CreatedAt:  crAt.In(jst),
			Text:       strings.Trim(strings.ReplaceAll(tweet.Text, "\n", ""), " "),
		}
		twr = append(twr, t)
	}

	//Tweetオブジェクトを返す
	return twr, nil
}
