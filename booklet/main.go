package main

import (
	"booklet/cli"
	"booklet/driver"
)

func main() {
	switch cli.Parse() {
	case cli.PullCmd.FullCommand():
		driver.Pull()
	case cli.CompileCmd.FullCommand():
		driver.Compile()
	}
}
