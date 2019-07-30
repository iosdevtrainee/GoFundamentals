package files

import (
	"fmt"
	"math/rand"
	"os"
)

func writeFile(contents string) (*string, error) {
	num := rand.ExpFloat64()
	filename := fmt.Sprintf("%f.txt", num)
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	_, err = file.WriteString(contents)
	if err != nil {
		return nil, err
	}
	return &filename, nil
}
