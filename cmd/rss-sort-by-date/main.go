package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/codesoap/erss"
	"github.com/codesoap/rss2"
	"github.com/docopt/docopt-go"
)

var usage = `
Sort all items of an RSS 2.0 file by pubDate, in increasing order.

Usage:
    rss-sort-by-date <file>
`

type conf struct {
	File string
}

type byDate []*rss2.Item

func (a byDate) Len() int      { return len(a) }
func (a byDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byDate) Less(i, j int) bool {
	return a[i].PubDate.Time.Before(a[j].PubDate.Time)
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
	rss, err := erss.ReadRSS(conf.File)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when reading the existing RSS:`, err.Error())
		os.Exit(3)
	}
	sort.Stable(byDate(rss.Channel.Items))
	if err = erss.WriteRSS(rss, conf.File); err != nil {
		fmt.Fprintln(os.Stderr, `Error when rendering:`, err.Error())
		os.Exit(5)
	}
}
