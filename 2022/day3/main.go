package main

import (
	"fmt"
	"strings"
	"aoc2022/common"
)

type rucksack struct {
	items string
	comp [2]string
}

func parseRucksack(data []string) []rucksack {
	ret := make([]rucksack,0)
	for _,line := range data {
		length := len(line)
		rs := rucksack{
			line,
			[2]string{line[0:length/2],line[length/2:]},
		}	
		ret = append(ret,rs)
	}	
	return ret
}

func (rs *rucksack) findSharedItem() byte {
	for _, item := range rs.comp[0] {
		if strings.ContainsRune(rs.comp[1],item) {
			return byte(item)
		}	
	}	
	return 0
}

func priority(item byte) int {
	if item >= 97 {
		return int(item - 96)
	}
	return int(item - 38)
}

func part1(data []string) int{
	rucksacks := parseRucksack(data)
	sum := 0
	for _, rs := range rucksacks {
		si := rs.findSharedItem() 
		sum += priority(si)
	}
	return sum
}

func findBadge(rucksacks []rucksack) byte {
	for _, item  := range rucksacks[0].items {
		if strings.ContainsRune(rucksacks[1].items,item) && strings.ContainsRune(rucksacks[2].items,item){
			return byte(item)
		}
	}
	return 0
}

func part2(data []string) int{
	rucksacks := parseRucksack(data)
	sum := 0
	for i:=0 ; i < len(data); i+=3{
		badge := findBadge(rucksacks[i:i+3])
		sum += priority(badge)
	}
	return sum
}

func main() {
	lines := common.Open("input.txt").Lines()
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
