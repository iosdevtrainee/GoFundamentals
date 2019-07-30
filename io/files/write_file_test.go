package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	expectedContents := "hello"
	filename, err := writeFile(expectedContents)
	if err != nil {
		t.Fatal(fmt.Sprintf("write file exited with error %s", err.Error()))
	}
	file, err := os.Open(*filename)
	if err != nil {
		t.Fatal(fmt.Sprintf("os.Open exited with error %s", err.Error()))
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(fmt.Sprintf("ioutil.ReadAll exited with error %s", err.Error()))
	}
	str := string(bytes[:])
	if str != expectedContents {
		t.Fatal(fmt.Sprintf("expected contents %s, received %s", expectedContents, str))
	}
}
