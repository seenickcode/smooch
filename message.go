package smooch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// SendTextMessage sends a message to a user using the 'appMaker' role.
func SendTextMessage(userID string, msg string, authToken string) (err error) {
	endpoint := fmt.Sprintf("%s%s/appusers/%s/messages", SmoochHostname, SmoochBaseEndpoint, userID)
	payload := struct {
		Text string `json:"text"`
		Role string `json:"role"`
		Type string `json:"type"`
	}{
		Text: msg,
		Role: "appMaker",
		Type: "text",
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}
	data, status, err := httpCallWithAuthToken("POST", endpoint, bytes.NewBuffer(payloadBytes), authToken)
	if err != nil {
		return
	}
	if status != http.StatusCreated {
		err = fmt.Errorf("posting Smooch conversation message failed with status %v: %v", status, string(data))
		return
	}
	return
}

// SendImageMessage sends an image with optional message to a user using the 'appMaker' role.
func SendImageMessage(userID string, optionalMsg string, mediaURL string, authToken string) (err error) {
	endpoint := fmt.Sprintf("%s%s/appusers/%s/messages", SmoochHostname, SmoochBaseEndpoint, userID)
	payload := struct {
		Text     string `json:"text,omitempty"`
		MediaURL string `json:"mediaUrl"`
		Role     string `json:"role"`
		Type     string `json:"type"`
	}{
		Text:     optionalMsg,
		MediaURL: mediaURL,
		Role:     "appMaker",
		Type:     "image",
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}
	data, status, err := httpCallWithAuthToken("POST", endpoint, bytes.NewBuffer(payloadBytes), authToken)
	if err != nil {
		return
	}
	if status != http.StatusCreated {
		err = fmt.Errorf("posting Smooch conversation message failed with status %v: %v", status, string(data))
		return
	}
	return
}

// SendReplyButtons sends a list of Reply Buttons to a user using the 'appMaker' role.
// Alternatively, Smooch "Button Shorthand" may be used instead via normal text (https://goo.gl/N0LsxE)
func SendReplyButtons(userID string, title string, buttons []*ReplyButton, authToken string) (err error) {
	lineParts := []string{}
	if len(title) > 0 {
		lineParts = append(lineParts, title)
	}
	buttonParts := []string{}
	for _, b := range buttons {
		buttonParts = append(buttonParts, fmt.Sprintf("%%[%s](reply:%s)", b.Label, b.Reply))
	}
	lineParts = append(lineParts, strings.Join(buttonParts, "\n"))

	return SendTextMessage(userID, strings.Join(lineParts, "\n"), authToken)
}
