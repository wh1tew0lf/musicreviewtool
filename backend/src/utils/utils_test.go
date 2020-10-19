package utils

import "testing"

func TestMessage(t *testing.T) {
	expectedStatus := true
	expectedMessage := "test"

	result := Message(expectedStatus, expectedMessage)

	actualStatus := result["status"]
	actualMessage := result["message"]

	if actualStatus != expectedStatus {
		t.Error("For status expected", expectedStatus, "got", actualStatus)
	}

	if actualMessage != expectedMessage {
		t.Error("For status expected", expectedMessage, "got", actualMessage)
	}
}
