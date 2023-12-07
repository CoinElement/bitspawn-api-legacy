/*

 */

package challenge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/bitspawngg/bitspawn-api/models"
)

type RequestBodyDataAdapter struct {
	Platform      string         `json:"platform"`
	PlayerId      string         `json:"player_id"`
	Scoring       string         `json:"scoring"`
	StartAt       int64          `json:"start_at"`
	EndAt         int64          `json:"end_at"`
	NumberOfGames int            `json:"number_of_games"`
	Mode          string         `json:"mode"`
	Weights       map[string]int `json:"weights"`
}

type AdapterData struct {
	Count    map[string]int `json:"count"`
	Score    map[string]int `json:"score"`
	Total    float64
	TeamName string `json:"teamName,omitempty"`
}

type ResponseBodyDataAdapter struct {
	Data    AdapterData `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

func (tsvc *ChallengeService) QueryWeightedScore(challengeRecord models.ChallengeRecord) (float64, map[string]*string, string, error) {
	challenge, err := tsvc.DB.GetChallengeData(challengeRecord.ChallengeId)
	if err != nil {
		return 0, nil, "", fmt.Errorf("Error in GetChallengeData of %s: %v", challengeRecord.ChallengeId, err)
	}
	scoring, err := stringToIntValue(challenge.Scoring)
	if err != nil {
		return 0, nil, "", fmt.Errorf("Error parsing scoring of %s: %v", challenge.ChallengeID, err)
	}
	platformMap := make(map[string]string)
	platformMap["BATTLENET"] = "battle"
	platformMap["PSN"] = "psn"
	platformMap["XBOX"] = "xbl"
	platformMap["NINTENDO"] = "nintendo"
	platformMap["ORIGIN"] = "origin"
	platformMap["STEAM"] = "steam"
	requestBodyStruct := RequestBodyDataAdapter{
		Platform:      platformMap[challengeRecord.Platform],
		PlayerId:      challengeRecord.PlayerId,
		StartAt:       challengeRecord.RegisterDate.Unix(),
		EndAt:         challengeRecord.ChallengeDeadline.Unix(),
		NumberOfGames: challenge.NumberOfGames,
		Mode:          challenge.GameMode,
		Weights:       scoring,
	}
	requestBody, err := json.Marshal(requestBodyStruct)
	if err != nil {
		return 0, nil, "", fmt.Errorf("Error in marshalling request body of Data Adapter API: %v", err)
	}
	client := &http.Client{}
	challengeType := strings.ToLower(challengeRecord.ChallengeType)
	req, err := http.NewRequest("POST", tsvc.conf.GetConfig().AdapterUrl+challengeType+"/weighted/score", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, nil, "", fmt.Errorf("Error making request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, "", fmt.Errorf("Error calling data adapter API: %v", err)
	}
	if resp.StatusCode != 200 {
		return 0, nil, "", fmt.Errorf("Data adapter API returns error status code: %d", resp.StatusCode)
	} else {
		defer resp.Body.Close()
	}

	var responseBody ResponseBodyDataAdapter
	bodyInBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyInBytes, &responseBody)
	if err != nil {
		return 0, nil, "", fmt.Errorf("Error in unmarshalling json payload from data adapter API: %v", err)
	}
	var score float64
	scoreMap := make(map[string]*string)
	if responseBody.Status != "success" {
		s := "0"
		for category := range scoring {
			scoreMap[category] = &s
		}
		score = 0
	} else {
		scoreMap = intToStringValue(responseBody.Data.Count)
		score = responseBody.Data.Total
	}
	return score, scoreMap, responseBody.Data.TeamName, nil
}

func stringToIntValue(mapToString map[string]*string) (map[string]int, error) {
	mapToInt := make(map[string]int)
	for key, value := range mapToString {
		if value == nil {
			mapToInt[key] = 0
		} else {
			v, err := strconv.Atoi(*value)
			if err != nil {
				return nil, err
			}
			mapToInt[key] = v
		}
	}
	return mapToInt, nil
}

func intToStringValue(mapToInt map[string]int) map[string]*string {
	mapToString := make(map[string]*string)
	for key, value := range mapToInt {
		v := strconv.Itoa(value)
		mapToString[key] = &v
	}
	return mapToString
}
