package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func intRange(x int64, y int64) []int64 {
	var out []int64
	for i := x; i <= y; i++ {
		out = append(out, i)
	}

	return out
}

func mapminmax(s map[int64]int64) (int64, int64) {
	min, max := s[0], s[0]

	for _, v := range s {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return min, max
}

func lastsplit(s string) int64 {
	var tmp []string

	if len(s) > 0 {
		tmp = strings.Split(strings.TrimSpace(s), " ")
		if len(tmp) <= 1 {
			return 2016
		}
	}

	out, _ := strconv.ParseInt(tmp[len(tmp)-1], 10, 64)
	return out
}

func main() {
	var aliveinyear = make(map[int64]int64)

	input, err := os.Open("presidents.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(input)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		for _, i := range intRange(lastsplit(record[1]), lastsplit(record[3])) {
			aliveinyear[i]++
		}
	}

	_, mostalivecount := mapminmax(aliveinyear)

	for year, count := range aliveinyear {
		if count == mostalivecount {
			fmt.Printf("%d:%d\n", year, count)
		}
	}
}
