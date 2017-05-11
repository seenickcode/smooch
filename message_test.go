package smooch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTexts(t *testing.T) {
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

	// send an image
	err = SendImageMessage(u1.ID,
		fmt.Sprintf("Test image %s", RandomStringWithLength(5)),
		"https://store-logos-us-east-1.s3.amazonaws.com/golang.png",
		s.NewAppAuthToken())
	assert.NoError(err)

	// send reply buttons
	btns1 := []*ReplyButton{
		{
			Label: RandomStringWithLength(5),
			Reply: RandomStringWithLength(5),
		},
		{
			Label: RandomStringWithLength(5),
			Reply: RandomStringWithLength(5),
		},
	}
	SendReplyButtons(u1.ID, RandomStringWithLength(5), btns1, s.NewAppAuthToken())
}
