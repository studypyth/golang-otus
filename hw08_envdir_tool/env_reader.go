package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := Environment{}
	files := make([]os.FileInfo, 0)
	items, err := ioutil.ReadDir(dir)

	for _, item := range items {
		if !item.IsDir() {
			files = append(files, item)
		}
	}
	for _, file := range files {
		b, err := ioutil.ReadFile(dir + "\\" + file.Name())
		if err != nil {
			panic(err)
		}

		value := strings.Split(string(b), "\r\n")
		line := value[0]

		if strings.ContainsRune(file.Name(), '=') {
			continue
		}

		switch {
		case file.Size() == 0:
			env[file.Name()] = EnvValue{Value: "", NeedRemove: true}
		default:
			fmt.Println(line)
			fmt.Println(file.Name())
			env[file.Name()] = EnvValue{Value: line, NeedRemove: false}
			fmt.Printf("%v", env)

		}
	}
	return env, err
}
