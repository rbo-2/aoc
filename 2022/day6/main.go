package main

import (
	"fmt"
	"aoc2022/common"
	"os"
	"strings"
)

func isMarker(stream string) bool {
	for i, ch := range stream {
		if strings.IndexByte(stream[i+1:], byte(ch)) != -1 {
			return false
		}
	}
	return true
}

func part1(line string) int {
	for i := range line {
		if i+4 > len(line) {
			break
		}
		if isMarker(line[i : i+4]) {
			fmt.Println(line[i : i+4])
			return i + 4
		}
	}
	return 0
}

func part2(line string) int {
	for i := range line {
		if i+4 > len(line) {
			break
		}
		if isMarker(line[i : i+14]) {
			fmt.Println(line[i : i+14])
			return i + 14
		}
	}
	return 0
}

func main() {
	lines := common.Open(os.Args[1]).Lines()
	fmt.Printf("part1: %d\n", part1(lines[0]))
	fmt.Printf("part2: %d\n", part2(lines[0]))
}
