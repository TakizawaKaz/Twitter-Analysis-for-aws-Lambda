package main

import (
	"github.com/dghubble/oauth1"
	"testing"
)



func TestSearchTwetterApi(t *testing.T) {


	config := oauth1.NewConfig(e.TwitterConsumerKey,  e.TwitterConsumerSecret)
	token := oauth1.NewToken(e.TwitterAccessToken, e.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	err := SearchTweet(httpClient,"#コロナウィルス")
	if err != nil{
		t.Fail()
	}
	//fmt.Println(string(result))
}
