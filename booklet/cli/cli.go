package cli

import (
	"booklet/pkg"
	"fmt"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()

	PullCmd = kingpin.Command("pull", "Pull the docker image.")

	CompileCmd   = kingpin.Command("compile", "Compile the book.")
	CompileDir   = CompileCmd.Arg("dir", "Directory of the booklet. Default is the current working directory.").Default(".").String()
	CompileFonts = CompileCmd.Flag("fonts", "The fonts directory. Default is '_fonts'.").Short('f').Default("_fonts").String()
)

func Parse() string {
	var err error
	_ = err

	kingpin.Version(pkg.Version)
	switch kingpin.Parse() {
	case PullCmd.FullCommand():
		if *Verbose {
			fmt.Println("Command:", PullCmd.FullCommand())
		}
		return PullCmd.FullCommand()
	case CompileCmd.FullCommand():
		*CompileDir, _ = filepath.Abs(*CompileDir)
		if !filepath.IsAbs(*CompileFonts) {
			*CompileFonts = filepath.Join(*CompileDir, *CompileFonts)
		}

		if *Verbose {
			fmt.Println("Command:", CompileCmd.FullCommand())
			tb := tablewriter.NewWriter(os.Stdout)
			tb.SetHeader([]string{"Arg", "Value"})
			tb.Append([]string{"Verbose", fmt.Sprintf("%t", *Verbose)})
			tb.Append([]string{"Dir", *CompileDir})
			tb.Append([]string{"Fonts", *CompileFonts})
			tb.Render()
			fmt.Println()
		}
		return PullCmd.FullCommand()
	}
	return ""
}
