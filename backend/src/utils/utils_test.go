package utils

import "testing"

func expectEqual(t *testing.T, field string, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Error("For", field, "expected", expected, "got", actual)
	}
}

func TestMessage(t *testing.T) {
	expectedStatus := true
	expectedMessage := "test"

	result := Message(expectedStatus, expectedMessage)

	actualStatus := result["status"]
	actualMessage := result["message"]

	expectEqual(t, "status", expectedStatus, actualStatus)
	expectEqual(t, "message", expectedMessage, actualMessage)
}
