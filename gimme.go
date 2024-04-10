package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	targetServiceName string = ""
)

type Generator struct {
	TargetServiceName string
	Endpoints         []Endpoint
}

type Endpoint struct {
	Name   string
	Path   string
	Method string
}

func main() {
	targetMDFile := "api.md"
	targetOutputFile := "generated_client.go"

	// Parse the MD file
	eps, err := parseLocalMDFile(targetMDFile)
	if err != nil {
		return
	}

	// Generate the code from the template provided
	generatedCode, err := generateCodeFromTemplate(eps)
	if err != nil {
		return
	}

	// Output to the target output file
	writeOutputToFile(targetOutputFile, generatedCode)
}

func parseLocalMDFile(fileName string) ([]Endpoint, error) {
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

		if strings.HasPrefix(line, "# ") {
			targetServiceName = strings.TrimLeft(line, "# ")
		}

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

func parseMDFileFromString() {
	// TODO: Parse the apimd file when it's retrieved as a string from a service response.
}

func writeOutputToFile(fileName string, content string) (bool, error) {
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return false, err
	}
	return true, nil
}

func generateCodeFromTemplate(endpoints []Endpoint) (string, error) {
	templateFile := "template.go.tmpl"
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, Generator{
		Endpoints:         endpoints,
		TargetServiceName: targetServiceName,
	})
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func fetchApiMDFromService(targetUrl string) (string, error) {
	resp, err := http.Get(targetUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to fetch api.md from remote url. StatusCode=%v", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
