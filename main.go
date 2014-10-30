package main

import (
	"flag"
	"fmt"
	"strings"
)

type ParsedCmd struct {
	Cmd  []string
	Args []string
}

func main() {
	// Parsing command line
	flag.Parse()
	args := flag.Args()
	cmd := &ParsedCmd{}
	if len(args) > 0 {
		cmd.Cmd = strings.Split(args[0], ":")
		cmd.Args = args[1:]
	}
	// -
	if cmd == nil {
		println("No command")
	} else {
		fmt.Printf("%v\n", *cmd)
	}
}
