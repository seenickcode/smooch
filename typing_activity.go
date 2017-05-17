package smooch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TypingActivityType string

const (
	TypingStart TypingActivityType = "typing:start"
	TypingStop  TypingActivityType = "typing:stop"
)

// ToggleTypingActivity shows/hides typing activity
func ToggleTypingActivity(userID string, activityType TypingActivityType, authToken string) (err error) {
	endpoint := fmt.Sprintf("%s%s/appusers/%s/conversation/activity", SmoochHostname, SmoochBaseEndpoint, userID)
	payload := struct {
		Role string             `json:"role"`
		Type TypingActivityType `json:"type"`
	}{
		Role: "appMaker",
		Type: activityType,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}
	data, status, err := httpCallWithAuthToken("POST", endpoint, bytes.NewBuffer(payloadBytes), authToken)
	if err != nil {
		return
	}
	if status != http.StatusOK {
		err = fmt.Errorf("setting Smooch conversation typing activity failed with status %v: %v", status, string(data))
		return
	}
	return
}
