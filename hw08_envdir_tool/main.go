package main

import (
	"os"
)

func main() {
	path := os.Args[1]
	commands := os.Args[2:]
	env, err := ReadDir(path)
	check(err)
	code := RunCmd(commands, env)
	os.Exit(code)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
