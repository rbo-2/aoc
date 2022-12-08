package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"aoc2022/common"
)

type dir struct {
	name     string
	parent   *dir
	children []*dir
	size     int
}

func changeDir(current *dir, dirName string) *dir {
	if dirName == ".." {
		return current.parent
	}
	for _, child := range current.children {
		if child.name == dirName {
			return child
		}
	}
	return nil
}

func updateParentsSize(curr *dir, size int) {
	if curr.parent == nil {
		return
	}
	curr.parent.size += size
	updateParentsSize(curr.parent, size)
}

func parseFileTree(lines []string) *dir {
	root := new(dir)
	root.name = "/"
	curr := &dir{
		children: []*dir{root},
	}
	root.parent = curr
	for _, line := range lines {
		if line[0] == '$' {
			args := strings.Split(line, " ")
			switch args[1] {
			case "ls":
				continue
			case "cd":
				curr = changeDir(curr, args[2])
				if curr == nil {
					panic("couldnt find directory")
				}
			}
		} else {
			// content of current dir
			content := strings.Split(line, " ")
			switch content[0] {
			case "dir":
				newDir := new(dir)
				newDir.name = content[1]
				newDir.parent = curr
				curr.children = append(curr.children, newDir)
			default:
				size, err := strconv.Atoi(content[0])
				if err != nil {
					panic("couldnt parse number!!")
				}
				curr.size += size
				updateParentsSize(curr, size)
			}
		}
	}
	return root
}

func dirSizes(d *dir, sizes []int) []int {
	sizes = append(sizes, d.size)
	for _, ch := range d.children {
		sizes = dirSizes(ch, sizes)
	}
	return sizes
}

func printTree(root *dir,indent int) {
	indentStr := strings.Repeat(" ",indent)
	fmt.Printf("%s - (dir) %s (%d) \n",indentStr,root.name,root.size)
	for _ , ch := range root.children {
		printTree(ch,indent+4)
	}
}

func part1(lines []string) int {
	root := parseFileTree(lines)
	sizes := make([]int, 0)
	sizes = dirSizes(root, sizes)
	size := 0
	for _, sz := range sizes {
		if sz < 100000 {
			size += sz
		}
	}
	return size
}

func part2(lines []string) int {
	root := parseFileTree(lines)
	sizes := make([]int, 0)
	sizes = dirSizes(root, sizes)
	unused := 70e6 - root.size
	sort.Ints(sizes)
	for _, sz := range sizes {
		if unused+sz >= 30e6 {
			return sz
		}
	}
	return 0
}

func part1Stack(lines []string) int {
	dirs := make([]string,0)
	sizes := make(map[string]int)
	for _, line := range lines {
		if line[0] == '$' {
			args := strings.Split(line, " ")
			switch args[1] {
			case "ls":
				continue
			case "cd":
				if args[2] == ".." {
					dirs = dirs[:len(dirs)-1]
					continue
				}
				dirs = append(dirs,args[2])
			}
		} else {
			// content of current dir
			content := strings.Split(line, " ")
			switch content[0] {
			case "dir":
				continue
			default:
				size, err := strconv.Atoi(content[0])
				if err != nil {
					panic("couldnt parse number!!")
				}
				for i := range dirs {
					sizes[strings.Join(dirs[:i+1],"/")] += size
				}
			}
		}
	}
	sz := 0
	for _,size := range sizes {
		if size < 100000 {
			sz+=size
		}
	}
	return sz
}

func main() {
	input := common.Open(os.Args[1]).Lines()
	fmt.Printf("part1: %d\n",part1(input))
	fmt.Printf("part2: %d\n",part1(input))
	fmt.Printf("part1(stack): %d\n",part1Stack(input))
}
