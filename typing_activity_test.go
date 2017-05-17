package smooch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypingActivity(t *testing.T) {
	assert := assert.New(t)
	s := NewTestSmooch(t)

	// create a user
	u1, err := PreCreateAppUser(RandomString(),
		fmt.Sprintf("User %s", RandomStringWithLength(5)),
		s.NewAppAuthToken())
	assert.NoError(err)
	assert.NotNil(u1)

	// send a message
	err = SendTextMessage(u1.ID,
		fmt.Sprintf("Test message %s", RandomStringWithLength(5)),
		s.NewAppAuthToken())
	assert.NoError(err)

	err = ToggleTypingActivity(u1.ID, TypingStart, s.NewAppAuthToken())
	assert.NoError(err)

	err = ToggleTypingActivity(u1.ID, TypingStop, s.NewAppAuthToken())
	assert.NoError(err)
}
