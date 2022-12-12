package main

import (
	"aoc2022/common"
	"fmt"
)

func findPos(grid [][]byte, source byte) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == source {
				return i, j
			}
		}
	}
	return -1, -1
}

func inBound(grid [][]byte, y, x int) bool {
	return len(grid) > y && y > -1 && x > -1 && len(grid[0]) > x
}

func stepsToDest(grid [][]byte, source []int, dest []int) int {
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	if source[0] == -1 {
		panic("couldnt find source pos")
	}
	grid[source[0]][source[1]] = 'a'
	grid[dest[0]][dest[1]] = 'z'
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	pathToSourceLen := make([][]int, len(grid))
	for i := range pathToSourceLen {
		pathToSourceLen[i] = make([]int, len(grid[0]))
	}
	pathToSourceLen[source[0]][source[1]] = 0
	visited[source[0]][source[1]] = true
	q := [][]int{}
	q = append(q, []int{source[0], source[1]})
	for len(q) != 0 {
		curPos := q[0]
		q = q[1:]
		for _, d := range dir {
			newY, newX := curPos[0]+d[0], curPos[1]+d[1]
			if !inBound(grid, newY, newX) || visited[newY][newX] {
				continue
			}
			cur := grid[curPos[0]][curPos[1]]
			adj := grid[newY][newX]
			if cur+1 == adj || adj <= cur {
				visited[newY][newX] = true
				q = append(q, []int{newY, newX})
				pathToSourceLen[newY][newX] += pathToSourceLen[curPos[0]][curPos[1]] + 1
			}
		}
	}
	return pathToSourceLen[dest[0]][dest[1]]
}

func part1(input []string) int {
	grid := make([][]byte, 0)
	for _, row := range input {
		rowList := make([]byte, 0)
		for _, elevation := range row {
			rowList = append(rowList, byte(elevation))
		}
		grid = append(grid, rowList)
	}
	y, x := findPos(grid, 'S')
	destY, destX := findPos(grid, 'E')
	stepsFromStoE := stepsToDest(grid, []int{y, x}, []int{destY, destX})
	return stepsFromStoE
}

func part2(input []string) int {
	grid := make([][]byte, 0)
	for _, row := range input {
		rowList := make([]byte, 0)
		for _, elevation := range row {
			rowList = append(rowList, byte(elevation))
		}
		grid = append(grid, rowList)
	}
	y, x := findPos(grid, 'S')
	destY, destX := findPos(grid, 'E')
	stepsFromStoE := stepsToDest(grid, []int{y, x}, []int{destY, destX})
	min := stepsFromStoE
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'a' {
				steps := stepsToDest(grid, []int{row, col}, []int{destY, destX})
				if steps < min && steps > 0 {
					min = steps
				}
			}
		}
	}
	return min
}

func main() {
	input := common.Open(common.Args(1)).Lines()
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
