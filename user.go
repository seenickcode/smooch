package smooch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID         string    `json:"_id"`
	UserID     string    `json:"userId"`
	GivenName  string    `json:"givenName"`
	SignedUpAt time.Time `json:"signedUpAt"`
	Properties struct {
		FavoriteFood string `json:"favoriteFood"`
	} `json:"properties"`
	ConversationStarted bool `json:"conversationStarted"`
	CredentialRequired  bool `json:"credentialRequired"`
}

func PreCreateAppUser(userID string, givenName string, authToken string) (user *User, err error) {
	endpoint := fmt.Sprintf("%s%s/appusers", SmoochHostname, SmoochBaseEndpoint)
	payload := struct {
		UserID    string `json:"userId"`
		GivenName string `json:"givenName"`
	}{
		UserID:    userID,
		GivenName: givenName,
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
		err = fmt.Errorf("pre-creating Smooch user failed with status %v: %v", status, string(data))
		return
	}
	// deserialize response
	respPayload := struct {
		AppUser *User `json:"appUser"`
	}{}
	err = json.Unmarshal(data, &respPayload)
	if err != nil {
		return
	}
	user = respPayload.AppUser
	return
}
