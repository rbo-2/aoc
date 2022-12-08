package common

import (
	"os"
	"strings"
)

type aocRawInput []byte

func Open(file string) aocRawInput {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return aocRawInput(data)
}

func (input aocRawInput) Lines() []string {
	return strings.Split(string(input[0:len(input)-1]), "\n")
}

func (input aocRawInput) SplitOn(sep string) []string {
	return strings.Split(string(input[0:len(input)-1]), sep)
}

func Args(i int) string {
	if len(os.Args) <= i {
		panic("need an argument!!!")
	}
	return os.Args[i]
}
