package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Endpoint struct {
	Name   string
	Path   string
	Method string
}

func main() {
	targetMDFile := "api.md"
	// targetOutputFile := "generated_client.go"

	// Parse the MD file
	eps, err := parseMDFile(targetMDFile)
	if err != nil {
		return
	}

	// Generate the code from the template provided

	// Output to the target output file
	fmt.Println(eps)

}

func parseMDFile(fileName string) ([]Endpoint, error) {
	var endpoints []Endpoint

	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	for i := 0; i < len(fileLines)-1; i++ {
		line := fileLines[i]
		if strings.HasPrefix(line, "### ") {
			endpoint := Endpoint{}
			innerIndex := i + 1
			for {
				// Base condition
				innerIndex++
				if innerIndex > len(fileLines)-1 {
					break
				}

				potentialInnerLine := fileLines[innerIndex]

				// Check if we've hit another endpoint, if so set outer index and break out
				if strings.HasPrefix(potentialInnerLine, "### ") {
					i = innerIndex - 1
					break
				}

				//Hacky af
				if strings.Contains(potentialInnerLine, "Name") {
					endpoint.Name = strings.TrimSpace(strings.Split(potentialInnerLine, ":")[1])
					continue
				}

				if strings.Contains(potentialInnerLine, "Path") {
					endpoint.Path = strings.TrimSpace(strings.Split(potentialInnerLine, ":")[1])
					continue
				}

				if strings.Contains(potentialInnerLine, "Method") {
					endpoint.Method = strings.TrimSpace(strings.Split(potentialInnerLine, ":")[1])
					continue
				}
			}
			endpoints = append(endpoints, endpoint)
		}
	}

	return endpoints, nil

}

func writeOutputToFile(fileName string, content string) (bool, error) {
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return false, err
	}
	return true, nil
}

func generateCodeFromTemplate(endpoints []Endpoint) (string, error) {
	return "nil", nil
}
