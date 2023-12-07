package tournament

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//
//	"github.com/bitspawngg/bitspawn-api/models"
//)
//
//type RequestBodyCreateMatchSchedule struct {
//	TournamentId string `json:"tournamentId"`
//}
//
//type ResponseBody struct {
//	Msg   string `json:"msg"`
//	Error string `json:"error"`
//}
//
//func (tsvc *TournamentService) CreateMatchSchedule(tournamentId string) error {
//	requestBodyStruct := RequestBodyCreateMatchSchedule{
//		TournamentId: tournamentId,
//	}
//	requestBody, err := json.Marshal(requestBodyStruct)
//	if err != nil {
//		return fmt.Errorf("Error in marshalling request body of Data Adapter API: %v", err)
//	}
//	client := &http.Client{}
//	req, err := http.NewRequest("POST", tsvc.organizerUrl+"/creatematchschedule", bytes.NewBuffer(requestBody))
//	resp, err := client.Do(req)
//	if err != nil {
//		return fmt.Errorf("Error calling CreateMatchSchedule API: %v", err)
//	}
//	if resp.StatusCode != 200 {
//		return fmt.Errorf("CreateMatchSchedule API returns error status code: %d", resp.StatusCode)
//	} else {
//		defer resp.Body.Close()
//	}
//
//	var responseBody ResponseBody
//	bodyInBytes, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(bodyInBytes, &responseBody)
//	if err != nil {
//		return fmt.Errorf("Error in unmarshalling json payload from CreateMatchSchedule API: %v", err)
//	}
//	if responseBody.Msg != "success" {
//		return fmt.Errorf("CreateMatchSchedule API returns error msg: %s", responseBody.Error)
//	}
//	return nil
//}
//
//type RequestBodyMockMatchSchedule struct {
//	Format string `json:"format"`
//	NTeams uint   `json:"nTeams"`
//}
//
//type ResponseBodyMockMatchSchedule struct {
//	Data  []models.Match `json:"data"`
//	Msg   string         `json:"msg"`
//	Error string         `json:"error"`
//}
//
//func (tsvc *TournamentService) MockMatchSchedule(format string, nTeams uint) ([]models.Match, error) {
//	requestBodyStruct := RequestBodyMockMatchSchedule{
//		Format: format,
//		NTeams: nTeams,
//	}
//	requestBody, err := json.Marshal(requestBodyStruct)
//	if err != nil {
//		return nil, fmt.Errorf("Error in marshalling request body of Data Adapter API: %v", err)
//	}
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", tsvc.organizerUrl+"/mockmatchschedule", bytes.NewBuffer(requestBody))
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf("Error calling MockMatchSchedule API: %v", err)
//	}
//	if resp.StatusCode != 200 {
//		return nil, fmt.Errorf("MockMatchSchedule API returns error status code: %d", resp.StatusCode)
//	} else {
//		defer resp.Body.Close()
//	}
//
//	var responseBody ResponseBodyMockMatchSchedule
//	bodyInBytes, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(bodyInBytes, &responseBody)
//	if err != nil {
//		return nil, fmt.Errorf("Error in unmarshalling json payload from MockMatchSchedule API: %v", err)
//	}
//	if responseBody.Msg != "success" {
//		return nil, fmt.Errorf("MockMatchSchedule API returns error msg: %s", responseBody.Error)
//	}
//
//	return responseBody.Data, nil
//}
//
//func (tsvc *TournamentService) PrepareManualMatches() error {
//	client := &http.Client{}
//	req, err := http.NewRequest("POST", tsvc.organizerUrl+"/preparemanualmatches", nil)
//	resp, err := client.Do(req)
//	if err != nil {
//		return fmt.Errorf("Error calling PrepareManualMatches API: %v", err)
//	}
//	if resp.StatusCode != 200 {
//		return fmt.Errorf("CreateMatchSchedule API returns error status code: %d", resp.StatusCode)
//	} else {
//		defer resp.Body.Close()
//	}
//
//	var responseBody ResponseBody
//	bodyInBytes, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(bodyInBytes, &responseBody)
//	if err != nil {
//		return fmt.Errorf("Error in unmarshalling json payload from PrepareManualMatches API: %v", err)
//	}
//	if responseBody.Msg != "success" {
//		return fmt.Errorf("PrepareManualMatches API returns error msg: %s", responseBody.Error)
//	}
//	return nil
//}
//
//func (tsvc *TournamentService) ReportWinners() error {
//	client := &http.Client{}
//	req, err := http.NewRequest("POST", tsvc.organizerUrl+"/reportwinners", nil)
//	resp, err := client.Do(req)
//	if err != nil {
//		return fmt.Errorf("Error calling reportwinners API: %v", err)
//	}
//	if resp.StatusCode != 200 {
//		return fmt.Errorf("reportwinners API returns error status code: %d", resp.StatusCode)
//	} else {
//		defer resp.Body.Close()
//	}
//
//	var responseBody ResponseBody
//	bodyInBytes, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(bodyInBytes, &responseBody)
//	if err != nil {
//		return fmt.Errorf("Error in unmarshalling json payload from reportwinners API: %v", err)
//	}
//	if responseBody.Msg != "success" {
//		return fmt.Errorf("reportwinners API returns error msg: %s", responseBody.Error)
//	}
//	return nil
//}
