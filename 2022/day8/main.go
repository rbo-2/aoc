package main

import (
	"aoc2022/common"
	"fmt"
)

type direction struct {
	x, y int
}

func inBound(row, col int, trees []string) bool {
	if row >= len(trees) || col >= len(trees[0]) {
		return false
	}
	if row < 0 || col < 0 {
		return false
	}
	return true
}

func checkDirection(tree byte, trees []string, row, col int, dir direction) bool {
	var walk func()
	result := false
	walk = func() {
		if !inBound(row, col, trees) {
			result = true
			return
		}
		if tree > trees[row][col] {
			row = row + dir.y
			col = col + dir.x
			walk()
		}
	}
	walk()
	return result
}

func viewDistance( tree byte, trees []string, row, col int, dir direction) int {
	treeCount := 0
	var walk func()
	walk = func() {
		if !inBound(row, col, trees) {
			return
		}
		treeCount++
		if tree > trees[row][col] {
			row = row + dir.y
			col = col + dir.x
			walk()
		}
	}
	walk()
	return treeCount
}

func isVisible(tree byte, trees []string, row, col int) bool {
	directions := []direction{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, dir := range directions {
		if checkDirection(tree, trees, row+dir.y, col+dir.x, dir) {
			return true
		}
	}
	return false
}

func part1(trees []string) int {
	count := 0
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			tree := trees[i][j]
			if isVisible(tree, trees, i, j) {
				count++
			}
		}
	}
	return len(trees)*2 + (len(trees[0])-2)*2 + count
}

func viewDistanceProduct(tree byte, trees []string, row, col int) int {
	directions := []direction{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	count := 1
	for _, dir := range directions {
		count *= viewDistance(tree, trees, row+dir.y, col+dir.x, dir)
	}
	return count
}

func part2(trees []string) int {
	max := 0
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			tree := trees[i][j]
			vd := viewDistanceProduct(tree, trees, i, j)
			if vd > max {
				max = vd
			}
		}
	}
	return max
}

func main() {
	input := common.Open(common.Args(1)).Lines()
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
