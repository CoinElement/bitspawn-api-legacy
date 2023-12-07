package models

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateMember(t *testing.T) {
	log := logrus.New()
	log.Out = os.Stdout
	db := getDB(t)
	tm := NewTeamMemberHandler(db)

	err := tm.CreateTeamMember("1", sub1, member, Apply)
	assert.NoError(t, err)
	err = tm.ChangeRole("1", sub1, manager)
	assert.NoError(t, err)
	updatedMember, err := tm.FindTeamMember("1", sub1)
	assert.NoError(t, err)
	assert.Equal(t, manager, updatedMember.Role)
	assert.Equal(t, updatedMember.Sub, updatedMember.User.Sub)
	err = tm.DeleteTeamMember("1", sub1)

	assert.NoError(t, err)
}
