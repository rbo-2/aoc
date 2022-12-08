package main

import (
	"fmt"
	"aoc2022/common"
	"os"
	"strconv"
	"strings"
)

type stack []byte

func (s *stack) push(b byte) {
	*s = append(*s, b)
}

func (s *stack) pop() byte {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *stack) reverse() {
	var rev stack
	for len(*s) != 0 {
		rev.push(s.pop())
	}
	*s = rev
}

type procedure struct {
	num, from, to int
}

func split(input string) []string {
	ret := make([]string,0)
	lo := 0
	hi := 3 
	for hi <= len(input) {
		ret = append(ret,input[lo:hi])
		hi+=4
		lo+=4
	}
	return ret
}

func parseStacks(input string) []stack {
	stackRows := strings.Split(input, "\n")
	stackNums := strings.TrimSpace(stackRows[len(stackRows)-1])
	l, err := strconv.Atoi(string(stackNums[len(stackNums)-1]))
	if err != nil {
		panic(err)
	}
	stacks := make([]stack, l)
	for _, v := range stackRows[0 : len(stackRows)-1] {
		items := split(v)
		for i, item := range items {
			if strings.TrimSpace(item) == "" {
				continue
			}
			stacks[i].push(byte(item[1]))
		}
	}
	for i := range stacks {
		stacks[i].reverse()
	}
	return stacks
}

func parseProcedures(input string) []procedure {
	ret := make([]procedure, 0)
	for _, row := range strings.Split(input, "\n") {
		words := strings.Split(row, " ")
		num, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(words[3])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(words[5])
		if err != nil {
			panic(err)
		}
		ret = append(ret, procedure{num, from-1, to-1})
	}
	return ret
}

func part1(input []string) string {
	stacks := parseStacks(input[0])
	procedures := parseProcedures(input[1])
	for _, p := range procedures {
		i := p.num
		for ;i > 0;i-- {
			stacks[p.to].push(stacks[p.from].pop())
		}
	}
	sb := strings.Builder{}
	for _,s := range stacks {
		sb.WriteByte(s.pop())
	}
	return sb.String()
}

func part2(input []string) string {
	stacks := parseStacks(input[0])
	procedures := parseProcedures(input[1])
	for _, p := range procedures {
		var tempStack stack
		for i := p.num ;i > 0;i-- {
			tempStack.push(stacks[p.from].pop())
		}
		for i := p.num ;i > 0;i-- {
			stacks[p.to].push((tempStack.pop()))
		}
	}
	sb := strings.Builder{}
	for _,s := range stacks {
		sb.WriteByte(s.pop())
	}
	return sb.String()
}

func main() {
	input := common.Open(os.Args[1]).SplitOn("\n\n")
	fmt.Printf("part1: %s\n",part1(input))
	fmt.Printf("part2: %s\n",part2(input))
}
