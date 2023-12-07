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

func TestFullFlowSocial(t *testing.T) {
	testSocialLinkAdd(t)
	testSocialLinkAdd(t)
	testSocialLinkList(t)
	testSocialLinkGet(t)
	testSocialLinkDelete(t)
}
func testSocialLinkAdd(t *testing.T) {
	url := fmt.Sprintf("http://localhost:%s%s", port, "/v2/socials")
	payload := `{
		"socialType" : "TWITTER",
		"socialID" : "social_id_1"
	}`
	resp, err := httpHelper(http.MethodPost, url, bytes.NewReader([]byte(payload)))
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusCreated, resp.StatusCode, "The final response did not match the expectation")

	body := resp.Body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err, "Could not reach server via url")
	respBody := &struct {
		Body models.SocialLink `json:"response"`
	}{}
	require.NoError(t, json.Unmarshal(data, respBody))
	require.Equal(t, models.Twitter, respBody.Body.SocialType, "response body should have the exact team Id")
	require.Equal(t, "social_id_1", respBody.Body.SocialID, "response body should have the exact team Id")
	require.Equal(t, "c025acff-5961-4973-ad3b-da803c828549", string(respBody.Body.UserID), "response body should have the exact team Id")
}
func testSocialLinkDelete(t *testing.T) {
	url := fmt.Sprintf("http://localhost:%s/v2/socials/%s", port, "TWITTER")

	resp, err := httpHelper(http.MethodDelete, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")
}
func testSocialLinkGet(t *testing.T) {
	url := fmt.Sprintf("http://localhost:%s/v2/socials/%s", port, "TWITTER")
	resp, err := httpHelper(http.MethodGet, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")

	body := resp.Body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err, "Could not reach server via url")
	respBody := &struct {
		Body models.SocialLink `json:"response"`
	}{}
	require.NoError(t, json.Unmarshal(data, respBody))
	require.Equal(t, models.Twitter, respBody.Body.SocialType, "response body should have the exact team Id")
	require.Equal(t, "social_id_1", respBody.Body.SocialID, "response body should have the exact team Id")
	require.Equal(t, "c025acff-5961-4973-ad3b-da803c828549", string(respBody.Body.UserID), "response body should have the exact team Id")

}

func testSocialLinkList(t *testing.T) {
	url := fmt.Sprintf("http://localhost:%s/v2/socials", port)
	resp, err := httpHelper(http.MethodGet, url, nil)
	require.NoError(t, err, "Could not reach server via url")
	require.Equal(t, http.StatusOK, resp.StatusCode, "The final response did not match the expectation")

	body := resp.Body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err, "Could not reach server via url")
	respBody := &struct {
		Body []*models.SocialLink `json:"response"`
	}{}
	require.NoError(t, json.Unmarshal(data, respBody))
	t.Logf("%+v\n", respBody)
	require.Equal(t, models.Twitter, respBody.Body[0].SocialType, "response body should have the exact team Id")
	require.Equal(t, "social_id_1", respBody.Body[0].SocialID, "response body should have the exact team Id")
	require.Equal(t, "c025acff-5961-4973-ad3b-da803c828549", string(respBody.Body[0].UserID), "response body should have the exact team Id")

}
