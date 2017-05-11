package smooch

import "fmt"

const (
	SmoochHostname     = "https://api.smooch.io"
	SmoochBaseEndpoint = "/v1"
)

type Smooch struct {
	appKeyID  string
	appSecret string
}

func NewSmooch(appKeyID string, appSecret string) (*Smooch, error) {
	if len(appKeyID) == 0 {
		return nil, fmt.Errorf("Smooch App Key ID is required.")
	}
	if len(appSecret) == 0 {
		return nil, fmt.Errorf("Smooch App Secret is required.")
	}
	return &Smooch{
		appKeyID:  appKeyID,
		appSecret: appSecret,
	}, nil
}
