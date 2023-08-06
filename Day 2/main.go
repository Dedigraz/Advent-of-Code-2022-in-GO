package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputstr := string(input)
	rounds := strings.Split(inputstr, "\r\n")
	sumscore := 0
	for _, r := range rounds {
		opp := strings.Split(r, " ")[0]
		me := strings.Split(r, " ")[1]

		sumscore += roundScore(opp, me)
	}

	fmt.Println(sumscore)
}

func roundScore(oppMove string, persMove string) int {
	outcome := 0
	elementScore := 0
	switch persMove {
	case "X":
		outcome += 0
		if oppMove == "A" {
			elementScore += 3
		}
		if oppMove == "B" {
			elementScore += 1
		}
		if oppMove == "C" {
			elementScore += 2
		}
	case "Y":
		outcome += 3
		if oppMove == "B" {
			elementScore += 2
		}
		if oppMove == "C" {
			elementScore += 3
		}
		if oppMove == "A" {
			elementScore += 1
		}
	case "Z":
		outcome += 6
		if oppMove == "C" {
			elementScore += 1
		}
		if oppMove == "A" {
			elementScore += 2
		}
		if oppMove == "B" {
			elementScore += 3
		}
	}

	return outcome + elementScore
}
