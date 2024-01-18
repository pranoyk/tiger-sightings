package utils

import (
	"encoding/base64"
	"strings"
	"testing"
	"time"
)

func TestDecodeCursor(t *testing.T) {
	// Test case 1: Empty cursor
	cursor := ""
	expectedTime := time.Now().UTC()
	expectedUUID := "00000000-0000-0000-0000-000000000000"

	res, uuid, err := DecodeCursor(cursor)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !res.After(expectedTime) {
		t.Errorf("Expected time %v, before %v", expectedTime, res)
	}
	if uuid != expectedUUID {
		t.Errorf("Expected UUID %v, got %v", expectedUUID, uuid)
	}

	// Test case 2: Valid encoded cursor
	encodedCursor := Encode(time.Now().UTC(), "example-uuid")
	res, uuid, err = DecodeCursor(encodedCursor)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if uuid != "example-uuid" {
		t.Errorf("Expected UUID %v, got %v", "example-uuid", uuid)
	}
}

func TestEncode(t *testing.T) {
	// Test case 1: Valid encoding
	lastSeen := time.Now().UTC()
	uuid := "example-uuid"

	encoded := Encode(lastSeen, uuid)
	decodedTime, decodedUUID, _ := DecodeCursor(encoded)

	if !decodedTime.Before(lastSeen) {
		t.Errorf("Expected decoded time %v, got %v", lastSeen, decodedTime)
	}
	if decodedUUID != "example-uuid" {
		t.Errorf("Expected UUID %v, got %v", "example-uuid", decodedUUID)
	}
}

func TestDecodeInvalidCursor(t *testing.T) {
	// Test case: Invalid encoded cursor
	invalidCursor := "invalid-cursor"
	_, _, err := DecodeCursor(invalidCursor)

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestDecodeInvalidBase64(t *testing.T) {
	// Test case: Invalid base64 encoded cursor
	invalidBase64Cursor := "invalid-base64-cursor"
	_, _, err := DecodeCursor(invalidBase64Cursor)

	if err == nil {
		t.Error("Expected an error, got nil")
	} else if !strings.Contains(err.Error(), "illegal base64 data") {
		t.Errorf("Expected error message containing 'illegal base64 data', got %v", err)
	}
}

func TestDecodeInvalidTimeFormat(t *testing.T) {
	// Test case: Invalid time format in the cursor
	invalidTimeFormatCursor := base64.StdEncoding.EncodeToString([]byte("invalid-time-format,example-uuid"))
	_, _, err := DecodeCursor(invalidTimeFormatCursor)

	if err == nil {
		t.Error("Expected an error, got nil")
	} else if !strings.Contains(err.Error(), "parsing time") {
		t.Errorf("Expected error message containing 'parsing time', got %v", err)
	}
}
