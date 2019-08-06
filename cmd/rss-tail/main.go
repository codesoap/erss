package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codesoap/erss"
	"github.com/codesoap/rss2"
	"github.com/docopt/docopt-go"
)

var usage = `
Remove all but the last n items form an existing RSS 2.0 file.

Usage:
    rss-tail [-n=<item_count>]
             [<infile>]
             [<outfile>]

Options:
    -n=<item_count>  Amount of items to keep [default: 30]
`

type conf struct {
	N       int
	Infile  string
	Outfile string
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
	rss, err := getRSS(&conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when reading the existing RSS:`, err.Error())
		os.Exit(3)
	}
	iStart := len(rss.Channel.Items) - conf.N
	if iStart < 0 {
		iStart = 0
	}
	rss.Channel.Items = rss.Channel.Items[iStart:]
	if err = erss.PrintOrWriteResult(rss, conf.Outfile); err != nil {
		fmt.Fprintln(os.Stderr, `Error when rendering:`, err.Error())
		os.Exit(5)
	}
}

// TODO: put into lib.go
func getRSS(conf *conf) (r *rss2.RSS, err error) {
	var input []byte
	if len(conf.Infile) == 0 {
		if input, err = ioutil.ReadAll(os.Stdin); err != nil {
			return
		}
	} else {
		if input, err = ioutil.ReadFile(conf.Infile); err != nil {
			return
		}
	}
	var rss rss2.RSS
	err = xml.Unmarshal(input, &rss)
	return &rss, err
}
