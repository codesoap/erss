package main

import (
	"fmt"
	"os"

	"github.com/codesoap/erss"
	"github.com/docopt/docopt-go"
)

var usage = `
Remove all but the last n items form an existing RSS 2.0 file.

Usage:
    rss-tail [-n=<item_count>] <file>

Options:
    -n=<item_count>  Amount of items to keep [default: 30]
`

type conf struct {
	N    int
	File string
}

func main() {
	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		// No need to print anything here, since ParseDoc() already does.
		os.Exit(1)
	}
	var conf conf
	err = opts.Bind(&conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when using arguments:`, err.Error())
		os.Exit(2)
	}
	rss, err := erss.GetRSS(conf.File)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when reading the existing RSS:`, err.Error())
		os.Exit(3)
	}
	iStart := len(rss.Channel.Items) - conf.N
	if iStart < 0 {
		iStart = 0
	}
	rss.Channel.Items = rss.Channel.Items[iStart:]
	if err = erss.WriteResult(rss, conf.File); err != nil {
		fmt.Fprintln(os.Stderr, `Error when rendering:`, err.Error())
		os.Exit(5)
	}
}
