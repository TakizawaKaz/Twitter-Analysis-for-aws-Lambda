package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

type Comprehend struct {
	c *comprehend.Comprehend
}

type Sentimental struct {
Positive   float64
Negative   float64
Mixid      float64
Natural    float64
}

func New(e Environment) Comprehend {

	// クレデンシャルの作成
	cred := credentials.NewStaticCredentials("[アクセスキーID]", "[シークレットアクセスキー]", "") // 最後の引数は[セッショントークン]
	// クレデンシャルとリージョンをセットしたコンフィグの作成
	region := "ap-northeast-1"
	conf := &aws.Config{
		Credentials: cred,
		Region:      &region,
	}
	// セッションの作成
	sess,err := session.NewSession(conf)
	ErrorCheck(err)

	svc := comprehend.New(sess)

	return Comprehend{ c: svc}
}

func (c * Comprehend) SentimentalAnalysis()(Sentimental, error)  {

}
