package common

import (
	"os"
	"strconv"
	"strings"
)

type AocRawInput []byte

func Open(file string) AocRawInput {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return AocRawInput(data)
}

func (input AocRawInput) Lines() []string {
	return strings.Split(string(input[0:len(input)-1]), "\n")
}

func (input AocRawInput) SplitOn(sep string) []string {
	return strings.Split(string(input[0:len(input)-1]), sep)
}

func Args(i int) string {
	if len(os.Args) <= i {
		panic("need an argument!!!")
	}
	return os.Args[i]
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}

func Int(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return num
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
