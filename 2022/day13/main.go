package main

import (
	"aoc2022/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type value struct {
	isNum     bool
	num       int
	items     []value
	isDivider bool
}

func isNum(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func parseValue(input string, pos int) (value, int) {
	if input[pos] == '[' {
		v := value{}
		v.isNum = false
		pos++
		for input[pos] != ']' {
			item, newPos := parseValue(input, pos)
			if input[newPos] == ',' {
				newPos++
			}
			pos = newPos
			v.items = append(v.items, item)
		}
		pos++
		return v, pos
	}
	i := pos
	for isNum(input[pos]) {
		pos++
	}

	n, err := strconv.Atoi(input[i:pos])
	if err != nil {
		panic(err)
	}
	return value{isNum: true, num: n}, pos
}

func S(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}

func compareValues(left, right value) int {
	if left.isNum {
		if right.isNum {
			return S(left.num - right.num)
		}
		return compareValues(value{isNum: false, items: []value{left}}, right)
	}
	if right.isNum {
		return compareValues(left, value{isNum: false, items: []value{right}})
	}
	ll, rl := len(left.items), len(right.items)
	for i := 0; i < ll && i < rl; i++ {
		if v := compareValues(left.items[i], right.items[i]); v != 0 {
			return v
		}
	}
	return S(ll - rl)
}

func part1(raw common.AocRawInput) int {
	input := raw.SplitOn("\n\n")
	sum := 0

	for i, p := range input {
		packets := strings.Split(p, "\n")
		left, _ := parseValue(packets[0], 0)
		right, _ := parseValue(packets[1], 0)
		if v := compareValues(left, right); v < 0 {
			sum += i + 1
		}
	}
	return sum
}

func part2(raw common.AocRawInput) int {
	input := raw.Lines()
	values := []value{}
	values = append(values, value{isNum: false, items: []value{{isNum: true, num: 2}}, isDivider: true})
	values = append(values, value{isNum: false, items: []value{{isNum: true, num: 6}}, isDivider: true})
	for _, p := range input {
		if len(p) == 0 {
			continue
		}
		v, _ := parseValue(p, 0)
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		if compareValues(values[i], values[j]) < 0 {
			return true
		}
		return false
	})
	res := 1
	for i, v := range values {
		if v.isDivider {
			fmt.Println(i)
			res *= (i + 1)
		}
	}
	return res
}

func main() {
	input := common.Open(common.Args(1))
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
