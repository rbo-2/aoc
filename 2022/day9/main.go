package main

import (
	"aoc2022/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

type rope struct {
	head, tail pos
	visited    []pos
}

type rope2 struct {
	knots   [10]pos
	visited []pos
}

func (r *rope2) updateHead(dir byte) {
	switch dir {
	case 'R':
		r.knots[0].x += 1
	case 'L':
		r.knots[0].x -= 1
	case 'U':
		r.knots[0].y += 1
	case 'D':
		r.knots[0].y -= 1
	}
}

func (r *rope2) updateKnots(head, tail int) {
	if head == 9 {
		return
	}
	dx := r.knots[head].x - r.knots[tail].x
	dy := r.knots[head].y - r.knots[tail].y
	if dy == 0 || dx == 0 {
		if math.Abs(float64(dy)) > 1 {
			r.knots[tail].y += sign(dy)
		} else if math.Abs(float64(dx)) > 1 {
			r.knots[tail].x += sign(dx)
		}
	} else if math.Abs(float64(dx)) != 1 || math.Abs(float64(dy)) != 1 {
		r.knots[tail].y += sign(dy)
		r.knots[tail].x += sign(dx)
	}
	r.updateKnots(head+1, tail+1)
}

func (r *rope2) addToVisited(p pos) {
	for _, ps := range r.visited {
		if p.x == ps.x && p.y == ps.y {
			return
		}
	}
	r.visited = append(r.visited, p)
}

func (r *rope) addToVisited(p pos) {
	for _, ps := range r.visited {
		if p.x == ps.x && p.y == ps.y {
			return
		}
	}
	r.visited = append(r.visited, p)
}

func sign(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	}
	return 0
}

func (r *rope) updateHead(dir byte) {
	switch dir {
	case 'R':
		r.head.x += 1
	case 'L':
		r.head.x -= 1
	case 'U':
		r.head.y += 1
	case 'D':
		r.head.y -= 1
	}
}

func (r *rope) updateTail() {
	dx := r.head.x - r.tail.x
	dy := r.head.y - r.tail.y
	if dy == 0 || dx == 0 {
		if math.Abs(float64(dy)) > 1 {
			r.tail.y += sign(dy)
		} else if math.Abs(float64(dx)) > 1 {
			r.tail.x += sign(dx)
		}
	} else if math.Abs(float64(dx)) != 1 || math.Abs(float64(dy)) != 1 {
		r.tail.y += sign(dy)
		r.tail.x += sign(dx)
	}
}

func part1(input []string) int {
	r := rope{
		head: struct{ x, y int }{0, 0},
		tail: struct{ x, y int }{0, 0},
	}
	for _, m := range input {
		words := strings.Split(m, " ")
		num, err := strconv.Atoi(string(words[1]))
		if err != nil {
			panic(err)
		}

		for n := 0; n < num; n++ {
			r.updateHead(words[0][0])
			r.updateTail()
			r.addToVisited(r.tail)
		}
	}
	return len(r.visited)
}

func part2(input []string) int {
	r := rope2{}
	for _, m := range input {
		words := strings.Split(m, " ")
		num, err := strconv.Atoi(string(words[1]))
		if err != nil {
			panic(err)
		}
		for n := 0; n < num; n++ {
			r.updateHead(words[0][0])
			r.updateKnots(0, 1)
			r.addToVisited(r.knots[9])
		}
	}
	return len(r.visited)
}

func main() {
	input := common.Open(common.Args(1)).Lines()
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
