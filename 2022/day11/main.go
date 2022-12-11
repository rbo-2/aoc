package main

import (
	"aoc2022/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operation struct {
		op byte
		n  int
	}
	test          int
	dest          [2]int
	inspectedItem int
}

func parseItems(items string) []int {
	res := make([]int, 0)
	for _, i := range strings.Split(items, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(i))
		if err != nil {
			panic(err)
		}
		res = append(res, n)
	}
	return res
}

func parseOperation(words string) (byte, int) {
	s := strings.Split(strings.TrimSpace(words), " ")
	op := s[len(s)-2]
	if s[len(s)-1] == "old" {
		return op[0], -1
	}
	n, err := strconv.Atoi(s[len(s)-1])
	if err != nil {
		panic(err)
	}

	return op[0], n
}

func parseTest(words string) int {
	s := strings.Split(strings.TrimSpace(words), " ")
	n, err := strconv.Atoi(s[len(s)-1])
	if err != nil {
		panic(err)
	}
	return n
}

func parseDest(words string) int {
	s := strings.Split(strings.TrimSpace(words), " ")
	n, err := strconv.Atoi(s[len(s)-1])
	if err != nil {
		panic(err)
	}
	return n
}

func part1(input []string) int {
	monkies := parseMonkies(input)
	round := 20
	for j := 0; j < round; j++ {
		for i, m := range monkies {
			for _, item := range m.items {
				monkies[i].inspectedItem++
				inspectingItem := item
				monkies[i].items = monkies[i].items[1:]
				op := m.operation
				if op.n == -1 {
					op.n = inspectingItem
				}

				switch op.op {
				case '*':
					inspectingItem = (inspectingItem * op.n) / 3
				case '+':
					inspectingItem = (inspectingItem + op.n) / 3
				case '-':
					inspectingItem = (inspectingItem - op.n) / 3
				case '/':
					inspectingItem = (inspectingItem / op.n) / 3
				}

				if inspectingItem%m.test == 0 {
					monkies[m.dest[0]].items = append(monkies[m.dest[0]].items, inspectingItem)

				} else {
					monkies[m.dest[1]].items = append(monkies[m.dest[1]].items, inspectingItem)
				}
			}
		}
	}
	sort.Slice(monkies, func(i, j int) bool {
		return monkies[i].inspectedItem < monkies[j].inspectedItem
	})

	return monkies[len(monkies)-1].inspectedItem * monkies[len(monkies)-2].inspectedItem
}

func part2(input []string) int {
	monkies := parseMonkies(input)
	N := 1
	for _, m := range monkies {
		N *= m.test
	}
	round := 10000
	for j := 0; j < round; j++ {
		for i, m := range monkies {
			for _, item := range m.items {
				monkies[i].inspectedItem++
				inspectingItem := item
				monkies[i].items = monkies[i].items[1:]
				op := m.operation
				if op.n == -1 {
					op.n = inspectingItem
				}

				switch op.op {
				case '*':
					inspectingItem = (inspectingItem * op.n) % N
				case '+':
					inspectingItem = (inspectingItem + op.n) % N
				case '-':
					inspectingItem = (inspectingItem - op.n) % N
				case '/':
					inspectingItem = (inspectingItem / op.n) % N
				}

				if inspectingItem%m.test == 0 {
					monkies[m.dest[0]].items = append(monkies[m.dest[0]].items, inspectingItem)
				} else {
					monkies[m.dest[1]].items = append(monkies[m.dest[1]].items, inspectingItem)
				}
			}
		}
	}
	sort.Slice(monkies, func(i, j int) bool {
		return monkies[i].inspectedItem < monkies[j].inspectedItem
	})

	return monkies[len(monkies)-1].inspectedItem * monkies[len(monkies)-2].inspectedItem
}

func parseMonkies(input []string) []monkey {
	monkies := make([]monkey, 0)
	for _, p := range input {
		m := monkey{}
		lines := strings.Split(p, "\n")
		for _, l := range lines {
			l = strings.TrimSpace(l)
			items := strings.Split(l, ":")
			switch l[0] {
			case 'S':
				m.items = parseItems(items[1])
			case 'O':
				m.operation.op, m.operation.n = parseOperation(items[1])
			case 'T':
				m.test = parseTest(items[1])
			case 'I':
				if l[3] == 't' {
					m.dest[0] = parseDest(items[1])
				} else {
					m.dest[1] = parseDest(items[1])
				}
			}
		}
		monkies = append(monkies, m)
	}

	return monkies
}

func main() {
	input := common.Open(common.Args(1)).SplitOn("\n\n")
	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
