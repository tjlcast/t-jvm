package main

import (
	"fmt"
	"t-jvm/classpath"
)

func main() {
	cmd := classpath.ParseCmd()
	if cmd.VersionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		classpath.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *classpath.Cmd) {
	fmt.Printf("classpath: %s class: %s args:%v\n",
		cmd.CpOption,
		cmd.Class,
		cmd.Args)
}
