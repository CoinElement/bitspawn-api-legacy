package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type formSMS struct {
	Message     string `json:"message"`
	PhoneNumber string `json:"phoneNumber"`
}

func SendViaAWS(requestBodyStruct formSMS) error {
	requestBody, err := json.Marshal(requestBodyStruct)
	if err != nil {
		return fmt.Errorf("Fail to marshall request body")
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://uhktabbfhb.execute-api.us-east-1.amazonaws.com/v1/", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("Fail to create a request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Fail to call AWS SMS API: %v", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("AWS SMS API returns error status code: %d", resp.StatusCode)
	} else {
		defer resp.Body.Close()
	}
	return nil
}

func SendViaMessageBird(requestBodyStruct formSMS) error {
	requestBody, err := json.Marshal(requestBodyStruct)
	if err != nil {
		return fmt.Errorf("Fail to marshall request body")
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://d52ydh0yvb.execute-api.us-east-1.amazonaws.com/Prod/send/", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("Fail to create a request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Fail to call Message Bird SMS API: %v", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Message Bird SMS API returns error status code: %d", resp.StatusCode)
	} else {
		defer resp.Body.Close()
	}
	return nil
}

func SendSMS(message string, phoneNumber string, provider string) error {
	switch {
	case provider == "messagebird":
		requestBodyStruct := formSMS{
			Message:     message,
			PhoneNumber: phoneNumber,
		}
		return SendViaMessageBird(requestBodyStruct)
	case provider == "aws":
		requestBodyStruct := formSMS{
			Message:     message,
			PhoneNumber: phoneNumber,
		}
		return SendViaAWS(requestBodyStruct)
	default:
		return fmt.Errorf("provider must be aws or messagebird")
	}
}
