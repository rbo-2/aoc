package main

import (
	"aoc2022/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(input []string) int { x := 1
	xInCycle := []int{}
	for _, i := range input {
		words := strings.Split(i, " ")
		if words[0] == "noop" {
			xInCycle = append(xInCycle, x)
		} else {
			n, err := strconv.Atoi(words[1])
			if err != nil {
				panic(err)
			}
			xInCycle = append(xInCycle, x)
			xInCycle = append(xInCycle, x)
			x += n
		}
	}
	sum := 0
	for i := 20; i < len(xInCycle); i += 40 {
		sum += i * xInCycle[i-1]
	}
	return sum
}

func part2(input []string) string {
	spriteX := 1
	crt := strings.Builder{}
	crtPos := 0
	for _, i := range input {
		words := strings.Split(i, " ")
		if words[0] == "noop" {
			if math.Abs(float64(spriteX-crtPos)) <= 1 {
				crt.WriteByte('#')
			} else {
				crt.WriteByte(' ')
			}
			crtPos = (crtPos+1) % 40
			if crtPos == 0 {
				crt.WriteByte('\n')
			}
		} else {
			n, err := strconv.Atoi(words[1])
			if err != nil {
				panic(err)
			}
			for i := 0; i < 2; i++ {
				if math.Abs(float64(spriteX-crtPos)) <= 1 {
					crt.WriteByte('#')
				} else {
					crt.WriteByte(' ')
				}
				crtPos = (crtPos+1) % 40
				if crtPos == 0 {
					crt.WriteByte('\n')
				}
			}
			spriteX += n
		}
	}
	return crt.String()
}

func main() {
	input := common.Open(common.Args(1)).Lines()
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Println(part2(input))
}
