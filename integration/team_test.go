package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bitspawngg/bitspawn-api/models"

	"github.com/stretchr/testify/require"
)

func TeamList(t *testing.T, teamID string) {
	url := fmt.Sprintf("http://localhost:%s%s", port, "/v1/teams")
	resp, err := httpHelper(http.MethodGet, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	body := resp.Body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err, "Could not reach server via url")
	respBody := &struct {
		Body []*models.TeamMachine `json:"response"`
	}{}
	require.NoError(t, json.Unmarshal(data, respBody))
	require.NotEmpty(t, teamID, respBody.Body[0].ID)
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")
}

func TestFullFlow(t *testing.T) {
	teamID := CreateTeam(t)
	ChangeTeam(t, teamID)
	TeamGet(t, teamID)
	TeamList(t, teamID)
	AddTeamMember(t, teamID)
	GetTeamMember(t, teamID, "4a231fa0-3a08-4607-8b30-9d0f27160b86")
	ChangeRole(t, teamID, "4a231fa0-3a08-4607-8b30-9d0f27160b86")
	Approve(t, teamID, "4a231fa0-3a08-4607-8b30-9d0f27160b86")
	GetTeamMembers(t, teamID)
	DeleteTeamMember(t, teamID, "4a231fa0-3a08-4607-8b30-9d0f27160b86")
	TeamDelete(t, teamID)
}

func CreateTeam(t *testing.T) string {
	url := fmt.Sprintf("http://localhost:%s%s", port, "/v1/teams")
	payload := `{
		"name" : "testteam1" ,
		"publicity" : "INVITE_ONLY",
		"genrePreferred" : "FPS"
	}`
	resp, err := httpHelper(http.MethodPost, url, bytes.NewReader([]byte(payload)))
	require.NoError(t, err, "Could not reach server via url")
	body := resp.Body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err, "Could not reach server via url")
	respBody := &struct {
		Body map[string]interface{} `json:"response"`
	}{}
	require.NoError(t, json.Unmarshal(data, respBody))
	teamID := respBody.Body["teamId"].(string)
	require.NotEmpty(t, teamID, "there is no teamId in the response")
	require.Equal(t, http.StatusCreated, resp.StatusCode, "The final response did not match the expectation")
	return teamID
}

func TeamDelete(t *testing.T, teamID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s", port, "/v1/teams", teamID)
	resp, err := httpHelper(http.MethodDelete, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")
}

func TeamGet(t *testing.T, teamID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s", port, "/v1/teams", teamID)
	resp, err := httpHelper(http.MethodGet, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")

	body := resp.Body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err, "Could not reach server via url")
	respBody := &struct {
		Body models.TeamMachine `json:"response"`
	}{}
	require.NoError(t, json.Unmarshal(data, respBody))
	require.Equal(t, teamID, respBody.Body.ID, "response body should have the exact team Id")
	require.Equal(t, "newvalue", respBody.Body.Name, "response body should have the exact team Id")
	require.Equal(t, "OPEN", string(respBody.Body.Publicity), "response body should have the exact team Id")
	require.Equal(t, "FPS", string(respBody.Body.GenrePreferred), "response body should have the exact team Id")

}

func GetTeamMember(t *testing.T, teamID, memberID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s/members/%s", port, "/v1/teams", teamID, memberID)
	resp, err := httpHelper(http.MethodGet, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")

	body := resp.Body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err, "Could not reach server via url")
	respBody := &struct {
		Body models.TeamMember `json:"response"`
	}{}
	require.NoError(t, json.Unmarshal(data, respBody))
	require.Equal(t, teamID, respBody.Body.TeamID, "response body should have the exact team Id")
	require.Equal(t, "4a231fa0-3a08-4607-8b30-9d0f27160b86", respBody.Body.Sub, "response body should have the exact team Id")

}

func GetTeamMembers(t *testing.T, teamID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s/members", port, "/v1/teams", teamID)
	resp, err := httpHelper(http.MethodGet, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")

	body := resp.Body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err, "Could not reach server via url")
	respBody := &struct {
		Body []*models.TeamMember `json:"response"`
	}{}
	require.NoError(t, json.Unmarshal(data, respBody))
	require.Equal(t, teamID, respBody.Body[0].TeamID, "response body should have the exact team Id")
	require.Equal(t, teamID, respBody.Body[1].TeamID, "response body should have the exact team Id")
	require.True(t, respBody.Body[0].Sub == "4a231fa0-3a08-4607-8b30-9d0f27160b86" || respBody.Body[1].Sub == "4a231fa0-3a08-4607-8b30-9d0f27160b86")
	t.Log("role", respBody.Body[0].Role, respBody.Body[1].Role)
	require.True(t, string(respBody.Body[0].Role) == "MANAGER" || string(respBody.Body[1].Role) == "MANAGER")
	require.Equal(t, "APPROVED", string(respBody.Body[1].Status))
	require.Equal(t, "APPROVED", string(respBody.Body[0].Status))

}

func AddTeamMember(t *testing.T, teamID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s/members", port, "/v1/teams", teamID)
	payload := `{
		"action" : "INVITE",
		"userId" : "4a231fa0-3a08-4607-8b30-9d0f27160b86"
	}`
	resp, err := httpHelper(http.MethodPost, url, bytes.NewReader([]byte(payload)))
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusCreated, resp.StatusCode, "The final response did not match the expectation")
}

func DeleteTeamMember(t *testing.T, teamID, memberID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s/members/%s", port, "/v1/teams", teamID, memberID)
	resp, err := httpHelper(http.MethodDelete, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")
}

func ChangeRole(t *testing.T, teamID, memberID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s/roles", port, "/v1/teams", teamID)
	payload := fmt.Sprintf(`{
		"role":"MANAGER",
		"userId":"%s"
		}`, memberID)
	resp, err := httpHelper(http.MethodPut, url, bytes.NewReader([]byte(payload)))
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")
}

func ChangeTeam(t *testing.T, teamID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s", port, "/v1/teams", teamID)
	payload := `{
		"name" : "newvalue" ,
		"publicity" : "OPEN",
		"genrePreferred" : "FPS"
	}`
	resp, err := httpHelper(http.MethodPut, url, bytes.NewReader([]byte(payload)))
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")
}

func Approve(t *testing.T, teamID, memberID string) {
	url := fmt.Sprintf("http://localhost:%s%s/%s/members/%s", port, "/v1/teams", teamID, memberID)
	payload := `{
		"action":"ACCEPT"
		}`
	resp, err := httpHelper(http.MethodPatch, url, bytes.NewReader([]byte(payload)))
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusNoContent, resp.StatusCode, "The final response did not match the expectation")
}

func TestNotFound(t *testing.T) {
	url := fmt.Sprintf("http://localhost:%s%s/%s", port, "/v1/teams", "8112")
	resp, err := httpHelper(http.MethodGet, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusNotFound, resp.StatusCode, "The final response did not match the expectation")
}
