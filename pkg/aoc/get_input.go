package aoc

import (
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

func getInputForDay(day int) (string, error) {
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

func GetInputForDay(day int) string {
	input, err := getInputForDay(day)
	if err != nil {
		log.Fatalf("failed to get input for day %d", day)
	}
	return strings.TrimSpace(input)
}

func GetInputLinesForDay(day int) []string {
	input, err := getInputForDay(day)
	if err != nil {
		log.Fatalf("failed to get input for day %d", day)
	}
	return strings.Split(strings.TrimSpace(input), "\n")
}
