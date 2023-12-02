package main

import (
	"bufio"
	"os"
)

// testing: do I need a comment
func LoadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result [] string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err:= scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}