package main

import (
	"fmt"
	"math/rand"
	"time"
)

var UsedMap = map[int]bool{}

var PerviousNumbers = make([]int, 0)

func main() {
	for len(UsedMap) < 90 {
		NextHousieNumber()
	}
}

func NextHousieNumber() {
	randumNum := CallRandumNumber()
	found := UsedMap[randumNum]
	for found {
		randumNum = CallRandumNumber()
		found = UsedMap[randumNum]
	}
	fmt.Println("Next Number is ", randumNum, " Previous Numbers ", PerviousNumbers)
	PerviousNumbers = append([]int{randumNum}, PerviousNumbers...)
	time.Sleep(7 * time.Second)
}

func CallRandumNumber() int {
	return rand.Intn(91) + 1
}
