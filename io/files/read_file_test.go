package files

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	filename := "test.txt"
	expectedOuput := "hello"
	output, err := readFile(filename)
	if err != nil {
		t.Fatal(fmt.Sprintf("reading file failed with error %s", err.Error()))
	}
	if *output != expectedOuput {
		t.Fatal(fmt.Sprintf("expected file to contain %s actual output is %s", expectedOuput, *output))
	}

}
