package main

import (
	// "bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	// "io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	input, err := os.ReadFile("input.txt")
	check(err)
	inputString := string(input)
	calories := strings.Split(inputString, "\r\n\r\n")
	var cummulatedC []int
	for i := 0; i < len(calories); i++ {
		cummulatedC = append(cummulatedC, cummulator(calories[i]))
	}

	sort.Ints(cummulatedC)
	top3 := cummulatedC[len(cummulatedC) -3:]

	max := 0

	for _, i := range top3 {
		max +=i
	}

	fmt.Print(max )
}

func cummulator(calorieList string) int{
	calories := strings.Split(calorieList, "\r\n")
	sum := 0
	for _, v := range calories {
		s , _ := strconv.Atoi(v)
		sum += s
	}
	return sum
}
