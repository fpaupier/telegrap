package handler

import (
	"testing"
)

func TestSanitizeStartCommand(t *testing.T) {
	var s = "/start yo dawg"
	var out = sanitize(s)
	var expected = " yo dawg"
	if expected != out {
		t.Errorf("expected %s, got %s", expected, out)
	}
}

func TestSanitizePunchCommand(t *testing.T) {
	var s = "/punch yo dawg"
	var out = sanitize(s)
	var expected = " yo dawg"
	if expected != out {
		t.Errorf("expected %s, got %s", expected, out)
	}
}

func TestSanitizeBotTagCommand(t *testing.T) {
	var s = "@RapGeniusBot yo dawg"
	var out = sanitize(s)
	var expected = " yo dawg"
	if expected != out {
		t.Errorf("expected %s, got %s", expected, out)
	}
}

func TestSanitizePunchAndBotTagCommand(t *testing.T) {
	var s = "/punch@RapGeniusBot yo dawg"
	var out = sanitize(s)
	var expected = " yo dawg"
	if expected != out {
		t.Errorf("expected %s, got %s", expected, out)
	}
}

func TestSanitizeCleanString(t *testing.T) {
	var s = "yo dawg"
	var out = sanitize(s)
	var expected = "yo dawg"
	if expected != out {
		t.Errorf("expected %s, got %s", expected, out)
	}
}
