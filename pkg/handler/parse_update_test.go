package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

var chat = Chat{Id: 1}

func TestParseUpdateMessageWithText(t *testing.T) {
	var msg = Message{
		Text: "hello world",
		Chat: chat,
	}

	var update = Update{
		UpdateId: 1,
		Message:  msg,
	}

	requestBody, err := json.Marshal(update)
	if err != nil {
		t.Errorf("Failed to marshal update in json, got %s", err.Error())
	}
	req := httptest.NewRequest("POST", "http://myTelegramWebHookHandler.com/secretToken", bytes.NewBuffer(requestBody))

	var updateToTest, errParse = parseTelegramRequest(req)
	if errParse != nil {
		t.Errorf("Expected a <nil> error, got %s", errParse.Error())
	}
	if *updateToTest != update {
		t.Errorf("Expected update %s, got %s", update, updateToTest)
	}

}

func TestParseUpdateMessageWithAudio(t *testing.T) {
	var audio = Audio{
		FileId:   "12345",
		Duration: 60,
	}
	var msg = Message{
		Chat:  chat,
		Audio: audio,
	}
	var update = Update{
		UpdateId: 1,
		Message:  msg,
	}
	requestBody, _ := json.Marshal(update)
	req := httptest.NewRequest("POST", "http://myTelegramWebHookHandler.com/secretToken", bytes.NewBuffer(requestBody))

	var updateToTest, err = parseTelegramRequest(req)

	if err != nil {
		t.Errorf("Expected a <nil> error, got %s", err.Error())
	}

	if *updateToTest != update {
		t.Errorf("Expected update %s, got %s", update, updateToTest)
	}

}

func TestParseUpdateMessageWithVoice(t *testing.T) {
	var voice = Voice{
		FileId:   "12345",
		Duration: 60,
	}
	var msg = Message{
		Chat:  chat,
		Voice: voice,
	}
	var update = Update{
		UpdateId: 1,
		Message:  msg,
	}
	requestBody, _ := json.Marshal(update)
	req := httptest.NewRequest("POST", "http://myTelegramWebHookHandler.com/secretToken", bytes.NewBuffer(requestBody))

	var updateToTest, err = parseTelegramRequest(req)

	if err != nil {
		t.Errorf("Expected a <nil> error, got %s", err.Error())
	}

	if *updateToTest != update {
		t.Errorf("Expected update %s, got %s", update, updateToTest)
	}

}

func TestParseUpdateMessageWithDocument(t *testing.T) {
	var doc = Document{
		FileId:   "12345",
		FileName: "sample.pdf",
	}
	var msg = Message{
		Chat:     chat,
		Document: doc,
	}
	var update = Update{
		UpdateId: 1,
		Message:  msg,
	}
	requestBody, _ := json.Marshal(update)
	req := httptest.NewRequest("POST", "http://myTelegramWebHookHandler.com/secretToken", bytes.NewBuffer(requestBody))

	var updateToTest, err = parseTelegramRequest(req)

	if err != nil {
		t.Errorf("Expected a <nil> error, got %s", err.Error())
	}

	if *updateToTest != update {
		t.Errorf("Expected update %s, got %s", update, updateToTest)
	}

}

func TestParseUpdateInvalid(t *testing.T) {
	var msg = map[string]string{
		"bipbip": "12345",
		"bopbop": "wrong input",
	}

	requestBody, _ := json.Marshal(msg)
	req := httptest.NewRequest("POST", "http://myTelegramWebHookHandler.com/secretToken", bytes.NewBuffer(requestBody))

	var _, err = parseTelegramRequest(req)

	if err == nil {
		t.Error("Expected an error, got <nil>")
	}
}
