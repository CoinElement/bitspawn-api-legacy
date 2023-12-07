package models

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	// connectionString is point to docker-compose
	// export CONNECTION_STRING= "host=localhost port=5432 user=postgres dbname=bitspawn password=123456"
	// two accounts has been created by initiate script
	sub1 string = "c025acff-5961-4973-ad3b-da803c828549"
	sub2 string = "4a231fa0-3a08-4607-8b30-9d0f27160b86"
)

func TestMain(m *testing.M) {
	os.Setenv("CONNECTION_STRING", "host=localhost port=5432 user=postgres dbname=bitspawn password=123456")
	code := m.Run()
	os.Exit(code)
}

// getDB create a db instance
func getDB(t *testing.T) *DB {
	log := logrus.New()
	log.Out = os.Stdout
	cnnString := os.Getenv("CONNECTION_STRING")
	t.Logf("%s", cnnString)
	db := NewDB("postgres", cnnString, log)
	assert.NoError(t, db.Connect())
	return db
}
func TestCreateTeam(t *testing.T) {
	db := getDB(t)
	th := NewTeamHandler(db)
	team1 := &TeamMachine{
		Name: "gold1",
	}
	team2 := &TeamMachine{
		Name: "gold2",
	}
	err := th.CreateTeamMachine(team1, sub1)
	assert.NoError(t, err)
	err = th.CreateTeamMachine(team2, sub1)
	assert.NoError(t, err)
	tm, err := th.FetchTeamMachineByID(team1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, tm)
	assert.Equal(t, "gold1", tm.Name)

	teams, err := th.FetchTeamMachinesByName("gold")
	assert.NoError(t, err)

	assert.True(t, team1.ID == teams[0].ID || team1.ID == teams[1].ID)
	err = th.DeleteTeamMachine(team1.ID)
	assert.NoError(t, err)
	err = th.DeleteTeamMachine(team2.ID)
	assert.NoError(t, err)
}

func TestTeamWithMembers(t *testing.T) {
	db := getDB(t)
	th := NewTeamHandler(db)
	tm := NewTeamMemberHandler(db)
	team := &TeamMachine{
		Name: "hi",
	}
	err := th.CreateTeamMachine(team, sub1)
	t.Logf("%+v", team)
	assert.NoError(t, err)
	err = tm.CreateTeamMember(team.ID, sub2, member, Invite)
	assert.NoError(t, err)
	fetchedTeam, err := th.FetchTeamMachineByID(team.ID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(fetchedTeam.Members))
	err = tm.DeleteTeamMember(team.ID, sub2)
	assert.NoError(t, err)
	fetchedTeam, err = th.FetchTeamMachineByID(team.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(fetchedTeam.Members))
	// owner
	err = tm.DeleteTeamMember(team.ID, sub1)
	assert.NoError(t, err)
	err = th.DeleteTeamMachine(team.ID)
	assert.NoError(t, err)
}

func TestChangeRole(t *testing.T) {
	db := getDB(t)
	th := NewTeamHandler(db)
	team := &TeamMachine{
		Name: "teamnumber1",
	}
	assert.NoError(t, th.CreateTeamMachine(team, sub1))
	assert.NoError(t, th.InviteToJoin(team.ID, sub2))
	fetchedTeam, err := th.FetchTeamMachineByID(team.ID)
	assert.NoError(t, err)
	assert.True(t, fetchedTeam.Members[0].Role == owner || fetchedTeam.Members[1].Role == owner)
	assert.True(t, fetchedTeam.Members[0].Role == member || fetchedTeam.Members[1].Role == member)
	assert.NoError(t, th.ChangeRole(team.ID, sub2, sub1, manager))
	fetchedTeam, err = th.FetchTeamMachineByID(team.ID)
	assert.NoError(t, err)
	assert.True(t, fetchedTeam.Members[0].Role == owner || fetchedTeam.Members[1].Role == owner)
	assert.True(t, fetchedTeam.Members[0].Role == manager || fetchedTeam.Members[1].Role == manager)
	assert.Equal(t, th.ChangeRole(team.ID, sub2, sub2, manager), ErrNotAllowed)
	assert.NoError(t, th.DeleteTeamMachine(team.ID))
}

func TestInvite(t *testing.T) {
	db := getDB(t)
	th := NewTeamHandler(db)
	team := &TeamMachine{
		Name:      "invite_team",
		Publicity: InviteOnly,
	}
	assert.NoError(t, th.CreateTeamMachine(team, sub1))
	assert.Error(t, th.AskToJoin(team.ID, sub2))
	assert.NoError(t, th.DeleteTeamMachine(team.ID))
}

func TestDuplicate(t *testing.T) {
	db := getDB(t)
	th := NewTeamHandler(db)
	team := &TeamMachine{
		Name:      "team_duplicate",
		Publicity: InviteOnly,
	}
	assert.NoError(t, th.CreateTeamMachine(team, sub1))
	team1 := &TeamMachine{
		Name:      "team_duplicate",
		Publicity: InviteOnly,
	}
	err := th.CreateTeamMachine(team1, sub1)
	assert.Equal(t, ErrDuplicatedEntity, convError(err))
	assert.NoError(t, th.DeleteTeamMachine(team.ID))

}

func TestUpdate(t *testing.T) {
	db := getDB(t)
	th := NewTeamHandler(db)
	teamOld := &TeamMachine{
		Name:           "team_old",
		Publicity:      InviteOnly,
		GenrePreferred: Sports,
	}
	assert.NoError(t, th.CreateTeamMachine(teamOld, sub1))
	teamNew := &TeamMachine{
		Name:           "team_new",
		Publicity:      Open,
		GenrePreferred: BattleRoyale,
	}
	assert.NoError(t, th.UpdateTeam(teamOld.ID, teamNew))
	team, err := th.FetchTeamMachineByID(teamOld.ID)
	assert.NoError(t, err)
	assert.Equal(t, team.Name, teamNew.Name)
	assert.Equal(t, team.Publicity, teamNew.Publicity)
	assert.Equal(t, team.GenrePreferred, teamNew.GenrePreferred)

	assert.NoError(t, th.DeleteTeamMachine(team.ID))

}
