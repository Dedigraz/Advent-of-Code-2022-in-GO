package main

import (
	"fmt"
	"os"
	"strings"
)

func main (){
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputstr := string(input)
	rounds := strings.Split(inputstr,"\r\n")
	sumscore :=0
	for _,r := range(rounds){
		opp := strings.Split(r, " ")[0]
		me := strings.Split(r, " ")[1]

		sumscore+= roundScore(opp,me)
	}

	fmt.Println(sumscore)
}

func roundScore( oppMove string, persMove string) int {
	elementScore := 0
	score := 0
	switch persMove {
	case "X":
		elementScore  += 1
		if oppMove == "A" {
			score += 3
		}
		if oppMove == "B" {
			score += 0
		}
		if oppMove == "C" {
			score += 6
		}
	case "Y":
		elementScore  += 2
		if oppMove == "B" {
			score += 3
		}
		if oppMove == "C" {
		score += 0
	}
		if oppMove == "A" {
		score += 6
		}
	case "Z":
		elementScore  += 3
		if oppMove == "C" {
		score += 3
	}
	if oppMove == "A" {
		score += 0
	}
			if oppMove == "B" {
			score += 6
		}
	}

	return elementScore + score
}