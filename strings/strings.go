package string

func reverse(input string) (reversedstr string) {
	size := len(input) - 1
	for index := range input {
		reversedstr += string(input[size-index])
	}
	return reversedstr
}

func numOfOcurrences(input string, testChar string) int {
	num := 0

	for _, char := range input {
		if string(char) == testChar {
			num++
		}
	}
	return num
}
