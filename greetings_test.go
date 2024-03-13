package greetings

import (
	"reflect"
	"testing"
)

func TestHello_WhenValidName_ThenGreetWithWelcomeMessage(t *testing.T) {
	input := "Alice"
	expected := "Hi Alice! Welcome!"

	got, err := Hello(input)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if got != expected {
		t.Errorf("Expected greeting %q, got %q", expected, got)
	}
}

func TestHello_WhenEmptyName_ThenReturnError(t *testing.T) {
	input := ""
	expectedError := "empty name"

	_, err := Hello(input)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != expectedError {
		t.Errorf("Expected error %q, got %q", expectedError, err.Error())
	}
}

func TestHellos_WhenMultipleNames_ThenReturnGreetingsForAll(t *testing.T) {
	input := []string{"Alice", "Bob"}
	expected := map[string]string{"Alice": "Hi Alice! Welcome!", "Bob": "Hi Bob! Welcome!"}

	got, err := Hellos(input)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected greetings %v, got %v", expected, got)
	}
}

func TestHellos_WhenEmptyListOfNames_ThenReturnError(t *testing.T) {
	input := []string{}
	expectedError := "empty list of names"

	_, err := Hellos(input)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != expectedError {
		t.Errorf("Expected error %q, got %q", expectedError, err.Error())
	}
}

func TestHellos_WhenNameWithError_ThenReturnError(t *testing.T) {
	input := []string{"", "Alice"}
	expectedError := "empty name"

	_, err := Hellos(input)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != expectedError {
		t.Errorf("Expected error %q, got %q", expectedError, err.Error())
	}
}
