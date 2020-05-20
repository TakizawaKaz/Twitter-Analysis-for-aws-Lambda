package utils

import (
	"fmt"
	"os"
)

const (
	AwsAccessKeyId        = "AWS_ACCESS_KEY_ID"
	AwsSecretAccessKey    = "AWS_SECRET_ACCESS_KEY"
	AwsRegion             = "AWS_REGION"
	AwsSessionToken       = "AWS_SESSION_TOKEN"
	TwitterConsumerKey    = "TWITTER_CONSUMER_KEY"
	TwitterConsumerSecret = "TWITTER_CONSUMER_SECRET"
	TwitterAccessToken    = "TWITTER_ACCESS_TOKEN"
	TwitterAccessSecret   = "TWITTER_ACCESS_SECRET"
	errMsg                = "get environment faile '%s' \n"
)

type Environment struct {
	AwsAccessKeyId        string
	AwsSecretAccessKey    string
	AwsSessionToken       string
	AwsRegion             string
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string
}

func (e *Environment) GetAccessKeys() (err error) {

	e.AwsAccessKeyId = os.Getenv(AwsAccessKeyId)
	e.AwsSecretAccessKey = os.Getenv(AwsSecretAccessKey)
	e.AwsSessionToken = os.Getenv(AwsSessionToken)
	e.AwsRegion = os.Getenv(AwsRegion)
	e.TwitterConsumerKey = os.Getenv(TwitterConsumerKey)
	e.TwitterConsumerSecret = os.Getenv(TwitterConsumerSecret)
	e.TwitterAccessToken = os.Getenv(TwitterAccessToken)
	e.TwitterAccessSecret = os.Getenv(TwitterAccessSecret)

	if e.AwsAccessKeyId == "" {
		return ErrorRaise(fmt.Sprintf(errMsg, AwsAccessKeyId))
	}
	if e.AwsSecretAccessKey == "" {
		return ErrorRaise(fmt.Sprintf(errMsg, AwsSecretAccessKey))
	}
	//if e.AwsSessionToken=="" {
	//	return ErrorRaise(fmt.Sprintf(errMsg, AwsSessionToken))
	//}
	if e.AwsRegion == "" {
		return ErrorRaise(fmt.Sprintf(errMsg, AwsRegion))
	}
	if e.TwitterConsumerKey == "" {
		return ErrorRaise(fmt.Sprintf(errMsg, TwitterConsumerKey))
	}
	if e.TwitterConsumerSecret == "" {
		return ErrorRaise(fmt.Sprintf(errMsg, TwitterConsumerSecret))
	}
	if e.TwitterAccessToken == "" {
		return ErrorRaise(fmt.Sprintf(errMsg, TwitterAccessToken))
	}
	if e.TwitterAccessSecret == "" {
		return ErrorRaise(fmt.Sprintf(errMsg, TwitterAccessSecret))
	}

	return nil
}
