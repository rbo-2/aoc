package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"aoc2022/common"
)


func parseCalories(data []string) ([]int,error) {
	ret := make([]int,2)
	for _,elf:=range data {
		cals := strings.Split(elf,"\n")
		arr := make([]int,0)
		sum := 0
		for _, cal := range cals {
			number,err:=strconv.Atoi(cal)
			if err != nil {
				return nil,err
			}
			arr = append(arr,number)
		}
		for _,v := range arr {
			sum += v
		}
		ret = append(ret,sum)
	}
	return ret,nil
}


func main() {
	lines := common.Open("input.txt").SplitOn("\n\n")	
	calories,err := parseCalories(lines)
	if err != nil {
		panic(err)
	}
	sort.Ints(calories)
	fmt.Printf("part1: %d\n",calories[len(calories)-1])
	fmt.Printf("part2: %d\n",calories[len(calories)-1]+calories[len(calories)-2]+calories[len(calories)-3])
}
