package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"aoc2022/common"
)

type pair struct {
	first, second struct {
		lo, hi int
	}
}

func isBetween(num, lo, hi int) bool {
	return num >= lo && num <= hi
}

func (p *pair) overlap() bool {
	if (p.first.lo <= p.second.lo && p.first.hi >= p.second.hi) ||
		(p.second.lo <= p.first.lo && p.second.hi >= p.first.hi) {
		return true
	}
	return false
}

func (p *pair) overlapAtAll() bool {
	if isBetween(p.second.hi, p.first.lo, p.first.hi) || isBetween(p.second.lo, p.first.lo, p.first.hi) ||
		isBetween(p.first.hi, p.second.lo, p.second.hi) || isBetween(p.first.lo, p.second.lo, p.second.hi) {
		return true
	}
	return false
}

func main() {
	input := common.Open(os.Args[1]).Lines()
	c1 := 0
	c2 := 0
	for _, line := range input {
		secs := strings.Split(line, ",")
		sec1 := strings.Split(secs[0], "-")
		sec2 := strings.Split(secs[1], "-")
		lo1, err := strconv.Atoi(string(sec1[0]))
		if err != nil {
			panic(err)
		}
		hi1, err := strconv.Atoi(string(sec1[1]))
		if err != nil {
			panic(err)
		}
		lo2, err := strconv.Atoi(string(sec2[0]))
		if err != nil {
			panic(err)
		}
		hi2, err := strconv.Atoi(string(sec2[1]))
		if err != nil {
			panic(err)
		}
		pair := pair{}
		pair.first.lo = lo1
		pair.first.hi = hi1
		pair.second.lo = lo2
		pair.second.hi = hi2
		if pair.overlap() {
			c1++
			c2++
		} else if pair.overlapAtAll() {
			c2++
		}
	}
	fmt.Printf("part1: %d\n", c1)
	fmt.Printf("part2: %d\n", c2)
}
