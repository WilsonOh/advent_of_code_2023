package aoc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fetchInputForDay(day int) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("failed to load .env file")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day), nil)
	if err != nil {
		return "", errors.New("failed to create request")
	}
	sessionCookie := os.Getenv("SESSION_COOKIE")
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", errors.New("failed to execute request")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}
	return string(body), nil
}

func getSampleFile(day int) (string, error) {
	fileName := fmt.Sprintf("./cmd/day%d/sample.txt", day)
	if fileExists(fileName) {
		log.Printf("Reading sample input file for day %d", day)
		input, err := os.ReadFile(fileName)
		return string(input), err
	}
	return "", errors.New("sample file does not exist. please create it first")
}

func getInputForDay(day int, isSample bool) (string, error) {
	if isSample {
		return getSampleFile(day)
	}
	fileName := fmt.Sprintf("./cmd/day%d/input.txt", day)
	if fileExists(fileName) {
		log.Printf("Input file for day %d already exists. Fetching input from file...", day)
		input, err := os.ReadFile(fileName)
		return string(input), err
	}

	log.Printf("Input file for day %d not found. Fetching input through a network request...", day)
	input, err := fetchInputForDay(day)
	if err != nil {
		return "", err
	}
	file, err := os.Create(fileName)
	if err != nil {
		return "", errors.Join(errors.New("failed to create input file"), err)
	}

	_, err = file.WriteString(input)
	if err != nil {
		return "", errors.Join(errors.New("failed to write to input file"), err)
	}
	return input, nil
}

func GetInputForDay(day int, isSample bool) string {
	input, err := getInputForDay(day, isSample)
	if err != nil {
		log.Fatalf("%v. Failed to get input for day %d.", err, day)
	}
	return strings.TrimSpace(input)
}

func GetInputScannerForDay(day int, isSample bool) *bufio.Scanner {
	input, err := getInputForDay(day, isSample)
	if err != nil {
		log.Fatalf("%v. Failed to get input for day %d.", err, day)
	}
	input = strings.TrimSpace(input)
	reader := strings.NewReader(input)
	return bufio.NewScanner(reader)
}

func GetInputLinesForDay(day int, isSample bool) []string {
	input, err := getInputForDay(day, isSample)
	if err != nil {
		log.Fatalf("%v. Failed to get input for day %d.", err, day)
	}
	return strings.Split(strings.TrimSpace(input), "\n")
}
