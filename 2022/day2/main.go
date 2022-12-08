package main

import (
	"fmt"
	"aoc2022/common"
)

const (
	rockPoint     = 1
	paperPoint    = 2
	scissorsPoint = 3
)

func part2(rounds []string) int {

	var resultAgainst = map[byte]struct {
		loss, win, draw int
	}{
		'A': {scissorsPoint, paperPoint, rockPoint}, // for A
		'B': {rockPoint, scissorsPoint, paperPoint}, // for B
		'C': {paperPoint, rockPoint, scissorsPoint}, // for C
	}

	score := 0

	for _, round := range rounds {
		roundScore := 0
		result := round[2]
		opp := round[0]
		switch result {
		case 'X':
			roundScore += resultAgainst[opp].loss
		case 'Y':
			roundScore += 3 + resultAgainst[opp].draw
		case 'Z':
			roundScore += 6 + resultAgainst[opp].win
		}
		score += roundScore
	}
	return score
}

type result struct{
	loss, win, draw byte 
}

var pairs = map[byte]result{
	'A': {'Z', 'Y', 'X'}, // for A
	'B': {'X', 'Z', 'Y'}, // for B
	'C': {'Y', 'X', 'Z'}, // for C
}

func resultPoint(r result,mine byte) int {
	switch mine{
	case r.loss:
		return 0
	case r.draw:
		return 3
	default:
		return 6
	}
}

func part1(rounds []string) int {

	score := 0

	for _, round := range rounds {
		roundScore := 0
		mine := round[2]
		opp := round[0]
		switch mine {
		case 'X':
			roundScore += rockPoint + resultPoint(pairs[opp],mine)
		case 'Y':
			roundScore += paperPoint +resultPoint(pairs[opp],mine)

		case 'Z':
			roundScore += scissorsPoint+resultPoint(pairs[opp],mine)
		}
		score += roundScore
	}
	return score
}

func main() {
	lines := common.Open("input.txt").Lines()
	fmt.Printf("part1: %d\n", part1(lines))
	fmt.Printf("part2: %d\n", part2(lines))
}
