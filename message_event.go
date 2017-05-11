package smooch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// MessageEvent represents a message sent via a Smooch Webook
type MessageEvent struct {
	Trigger string `json:"trigger"`
	App     struct {
		ID string `json:"_id"`
	} `json:"app"`
	Messages []struct {
		ID       string  `json:"_id"`
		Type     string  `json:"type"`
		Text     string  `json:"text"`
		Role     string  `json:"role"`
		AuthorID string  `json:"authorId"`
		Name     string  `json:"name"`
		Received float64 `json:"received"`
		Source   struct {
			Type string `json:"type"`
		} `json:"source"`
	} `json:"messages"`
	AppUser struct {
		ID        string `json:"_id"`
		UserID    string `json:"userId"`
		GivenName string `json:"givenName"`
		Surname   string `json:"surname"`
		// TODO, persist email  Email      string `json:"email"`
		Properties struct {
		} `json:"properties"`
		SignedUpAt time.Time `json:"signedUpAt"`
		Clients    []struct {
			Active   bool      `json:"active"`
			ID       string    `json:"id"`
			LastSeen time.Time `json:"lastSeen"`
			Platform string    `json:"platform"`
		} `json:"clients"`
	} `json:"appUser"`
}

func NewMessageEventFromRequest(r *http.Request) (m *MessageEvent, err error) {
	body, err := ioutil.ReadAll(r.Body)
	var obj MessageEvent
	if len(string(body)) > 0 {
		if err := json.Unmarshal(body, &obj); err != nil {
			return nil, err
		}
		m = &obj
	}
	return
}

func (m *MessageEvent) Text() string {
	if len(m.Messages) == 1 {
		return m.Messages[0].Text
	}
	parts := []string{}
	for _, msg := range m.Messages {
		parts = append(parts, msg.Text)
	}
	return strings.Join(parts, "\n")
}

func (m MessageEvent) UserID() string {
	return m.AppUser.ID
}
