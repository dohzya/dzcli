package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type ParsedCmd struct {
	Cmd  []string
	Args []string
}

func (p *ParsedCmd) CmdString() string {
	return strings.Join(p.Cmd, ":")
}

type Cmd struct {
	Name    string
	Func    func(...string)
	SubCmds map[string]*Cmd
}

func (c *Cmd) Add(cmd *Cmd) {
	c.SubCmds[cmd.Name] = cmd
}

func CreateCmdRoot() *Cmd {
	return &Cmd{SubCmds: make(map[string]*Cmd)}
}
func CreateCmd(name string, f func(...string)) *Cmd {
	return &Cmd{Name: name, Func: f, SubCmds: make(map[string]*Cmd)}
}
func CreateCmdNS(name string) *Cmd {
	return &Cmd{Name: name, SubCmds: make(map[string]*Cmd)}
}

func (c *Cmd) Get(arg []string) *Cmd {
	if len(arg) == 0 {
		return c
	} else {
		cmd := c.SubCmds[arg[0]]
		return cmd.Get(arg[1:])
	}
}

func (c *Cmd) Call(parsed *ParsedCmd) error {
	cmd := c.Get(parsed.Cmd)
	if cmd == nil {
		return fmt.Errorf("Command not found: %v", parsed.CmdString())
	} else if cmd.Func == nil {
		if len(parsed.Cmd) == 0 {
			return fmt.Errorf("Missing command")
		} else {
			return fmt.Errorf("%v is not a valid command", parsed.CmdString())
		}
	} else {
		cmd.Func(parsed.Args...)
		return nil
	}
}

var MainCmd = CreateCmdRoot()

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

	if err := MainCmd.Call(cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}
