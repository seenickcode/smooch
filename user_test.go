package smooch

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)
	s := NewTestSmooch(t)

	// create a user
	u1, err := PreCreateAppUser(RandomString(),
		fmt.Sprintf("User %s", RandomStringWithLength(5)),
		s.NewAppAuthToken())
	assert.NoError(err)
	assert.NotNil(u1)
}

func NewTestSmooch(t *testing.T) *Smooch {
	appKeyID := os.Getenv("SMOOCH_APP_KEY_ID")
	if len(appKeyID) == 0 {
		t.Fatalf("please set env var SMOOCH_APP_KEY_ID to run tests")
	}
	appSecret := os.Getenv("SMOOCH_APP_SECRET")
	if len(appSecret) == 0 {
		t.Fatalf("please set env var SMOOCH_APP_SECRET to run tests")
	}
	s, err := New(appKeyID, appSecret)
	if err != nil {
		t.Fatal(err)
	}
	return s
}
