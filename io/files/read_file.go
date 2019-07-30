package files

import (
	"io/ioutil"
	"os"
)

func readFile(filename string) (*string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	str := string(contents[:])
	return &str, nil
}
