package main

import (
	"fmt"
	"strings"
)

func add(args ...string) {
	if len(args) == 0 {
		fmt.Printf("Missing variables to add\n")
		return
	}
	toAdd := make(map[string]string)
	err := false
	for _, arg := range args {
		if split := strings.SplitN(arg, "=", 2); len(split) == 2 {
			name := split[0]
			value := split[1]
			toAdd[name] = value
		} else {
			fmt.Printf("Missing value: %v\n", arg)
			err = true
		}
	}
	if err {
		return
	}
	for name, value := range toAdd {
		fmt.Printf("Add %v (with value '%v') to environment\n", name, value)
	}
}

func init() {
	envCmd := CreateCmdNS("env")
	envCmd.Add(CreateCmd("add", add))
	MainCmd.Add(envCmd)
}
