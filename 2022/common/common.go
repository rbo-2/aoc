package common

import (
	"os"
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
