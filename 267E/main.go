package main

import (
	"fmt"
	"math"
	"os"
)

func getremainder(i int) (remainder int) {
	remainder = int(math.Mod(float64(i), float64(10)))

	for remainder > 10 {
		remainder = getremainder(int(remainder))
	}

	return remainder
}

func getordinal(i int) (o string) {
	var remainder int

	remainder = getremainder(i)

	switch remainder {
	case 1:
		return fmt.Sprintf("%dst", remainder)
	case 2:
		return fmt.Sprintf("%dnd", remainder)
	case 3:
		return fmt.Sprintf("%drd", remainder)
	default:
		return fmt.Sprintf("%dth", remainder)
	}
}

func notmydog(mydog int) (otherdogs []int, err error) {
	if mydog < 0 || mydog > 100 {
		fmt.Errorf("Dog must have placed from 0-100: Input was %i\n", mydog)
		return nil, err
	}

	for i := 0; i <= 100; i++ {
		if i != mydog {
			otherdogs = append(otherdogs, i)
		}
	}
	return otherdogs, nil
}

// Take integer input and calculate numbers from 0-100 that are !input and output in ordinal notation (1st, 2nd, 3rd)
func main() {
	fmt.Println("Reddit Daily Exercise 257E")
	mydog := 1

	otherdogs, err := notmydog(mydog)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	for _, v := range otherdogs {
		fmt.Printf("%v\n", getordinal(v))
	}
}
