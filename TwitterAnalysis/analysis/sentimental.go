package analysis

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
)

type Client struct {
	ComprehendClient *comprehend.Comprehend
	DetectText       []string
	LanguageCode     string
}

type SentimentalData struct {
	Sentiment string  `json:"sentiment"`
	Positive  float64 `json:"positive"`
	Negative  float64 `json:"negative"`
	Mixed     float64 `json:"mixed"`
	Neutral   float64 `json:"neutral"`
}

const LanguageCode = "ja"

func NewAnalysis(svc *comprehend.Comprehend, text []string) Client {
	return Client{
		ComprehendClient: svc,
		DetectText:       text,
		LanguageCode:     LanguageCode,
	}
}

func (c *Client) SentimentalAnalysis() ([]SentimentalData, error) {

	analysisData := make([]SentimentalData, 0, len(c.DetectText))

	for _, t := range c.DetectText {
		resp, err := c.ComprehendClient.DetectSentiment(
			&comprehend.DetectSentimentInput{
				LanguageCode: &c.LanguageCode,
				Text:         &t,
			})

		if err != nil {
			return nil, err
		}

		analysisData = append(analysisData,
			SentimentalData{
				Sentiment: *resp.Sentiment,
				Positive:  *resp.SentimentScore.Positive,
				Negative:  *resp.SentimentScore.Negative,
				Mixed:     *resp.SentimentScore.Mixed,
				Neutral:   *resp.SentimentScore.Neutral,
			})
	}

	return analysisData, nil
}
