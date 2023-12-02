# Advent of Code 2023
My solutions for Advent of Code 2023 in Golang, for the purpose of learning Golang

# Project Structure
```shell
.
├── cmd
│  ├── day1
│  │  ├── input.txt
│  │  ├── main.go
│  │  ├── sample.txt
│  └── day2
│     ├── input.txt
│     ├── main.go
│     └── sample.txt
├── go.mod
├── go.sum
└── pkg
   └── aoc
      └── get_input.go
```
# Addtional Information
The `aoc` package defined in `advent_of_code_2023/pkg/aoc` contains various helper functions, most notably `func GetInputForDay(day int) string`, which uses the
session cookie provided in `.env` to make a HTTP request to `https://adventofcode.com/2023/day/{day}/input` to retrieve the input for the day. The function
also saves the retrieved input to `cmd/day{day}/input.txt` so that subsequent calls to `GetInputForDay` file read from that file instead of making a network call.
