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

func getordinal(i int) string {
	var remainder int

	switch i {
	case 11:
		return fmt.Sprintf("%dth", i)
	case 12:
		return fmt.Sprintf("%dth", i)
	case 13:
		return fmt.Sprintf("%dth", i)
	}

	remainder = getremainder(i)

	switch remainder {
	case 1:
		return fmt.Sprintf("%dst", i)
	case 2:
		return fmt.Sprintf("%dnd", i)
	case 3:
		return fmt.Sprintf("%drd", i)
	default:
		return fmt.Sprintf("%dth", i)
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
