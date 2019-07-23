package string

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := "hello"
	expectedOutput := "olleh"
	result := reverse(input)
	if result != expectedOutput {
		t.Fatal(fmt.Sprintf("expected %s and receieved %s", expectedOutput, result))
	}
}

func TestNumberOfOccurrences(t *testing.T) {
	input := "hellllo"
	testLetter := "l"
	expectedResult := 4
	result := numOfOcurrences(input, testLetter)
	if expectedResult != result {
		t.Fatal(fmt.Sprintf("expected %d and receieved %d", expectedResult, result))
	}
}
