package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSocialLink(t *testing.T) {
	sh := NewSocialHandler(getDB(t))
	assert.NotNil(t, sh, "handler shouldn't be nil")
	err := sh.CreateSocialLink(&SocialLink{"mytestID", Twitter, "socialid_1"})
	assert.NoError(t, err, "no error expected")
	sl, err := sh.FetchSocialLink("mytestID", Twitter)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, sl.UserID, "mytestID")
	assert.Equal(t, sl.SocialID, "socialid_1")
	assert.Equal(t, sl.SocialType, Twitter)
	err = sh.DeleteSocialLink("mytestID", Twitter)
	assert.NoError(t, err, "no error expected")
	_, err = sh.FetchSocialLink("mytestID", Twitter)
	assert.Error(t, ErrRecordNotFound, err, "no error expected")
}
