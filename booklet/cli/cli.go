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
	Dir     = kingpin.Arg("dir", "Directory of the booklet. Default is the current working directory.").Default(".").String()
	Fonts   = kingpin.Flag("fonts", "The fonts directory. Default is '_fonts'.").Short('f').Default("_fonts").String()
)

func Parse() {
	var err error
	_ = err

	kingpin.Version(pkg.Version)
	kingpin.Parse()

	*Dir, _ = filepath.Abs(*Dir)

	if !filepath.IsAbs(*Fonts) {
		*Fonts = filepath.Join(*Dir, *Fonts)
	}

	if *Verbose {
		tb := tablewriter.NewWriter(os.Stdout)
		tb.SetHeader([]string{"Arg", "Value"})
		tb.Append([]string{"Verbose", fmt.Sprintf("%t", *Verbose)})
		tb.Append([]string{"Dir", *Dir})
		tb.Append([]string{"Fonts", *Fonts})
		tb.Render()
		fmt.Println()
	}
}
