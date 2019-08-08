package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/codesoap/erss"
	"github.com/codesoap/rss2"
	"github.com/docopt/docopt-go"
)

// TODO: Think about enabling multiple --category parameters.

var usage = `
Copy items form one RSS 2.0 file to another, respecting optional filters.
<title>, <category> and <guid> are interpreted as regular expressions.

Usage:
    rss-copy-items [--title=<title>]
                   [--category=<category>]
                   [--guid=<guid>]
                   [--after=<date>]
                   [--before=<date>]
                   <source>
                   <target>
`

type conf struct {
	Title    string
	Category string
	Guid     string
	After    string
	Before   string
	Source   string
	Target   string
}

var reTitle, reCategory, reGUID *regexp.Regexp

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
	err = setRegexps(&conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when using given regex:`, err.Error())
		os.Exit(2)
	}
	items, err := getFilteredItems(&conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when getting the selected items:`, err.Error())
		os.Exit(3)
	}
	target, err := erss.ReadRSS(conf.Target)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when reading the target RSS:`, err.Error())
		os.Exit(4)
	}
	target.Channel.Items = append(target.Channel.Items, items...)
	if err = erss.WriteResult(target, conf.Target); err != nil {
		fmt.Fprintln(os.Stderr, `Error when rendering:`, err.Error())
		os.Exit(5)
	}
}

func setRegexps(conf *conf) (err error) {
	if reTitle, err = regexp.Compile(conf.Title); err != nil {
		return
	}
	if reCategory, err = regexp.Compile(conf.Category); err != nil {
		return
	}
	reGUID, err = regexp.Compile(conf.Guid)
	return
}

func getFilteredItems(conf *conf) (itemsFiltered []*rss2.Item, err error) {
	source, err := erss.ReadRSS(conf.Source)
	if err != nil {
		return nil, fmt.Errorf(`error when reading source rss: %s`, err.Error())
	}
	itemsSource := source.Channel.Items
	for _, item := range itemsSource {
		var matches bool
		if matches, err = matchesFilter(item, conf); err != nil {
			return
		}
		if matches {
			itemsFiltered = append(itemsFiltered, item)
		}
	}
	return
}

func matchesFilter(item *rss2.Item, conf *conf) (matches bool, err error) {
	if !reTitle.Match([]byte(item.Title)) {
		return false, nil
	}
	if len(conf.Category) > 0 {
		matches = false
		for _, category := range item.Categories {
			matches = reCategory.Match([]byte(category.Value))
			if matches {
				break
			}
		}
		if !matches {
			return false, nil
		}
	}
	if len(conf.Guid) > 0 && (item.GUID == nil ||
		!reGUID.Match([]byte(item.GUID.Value))) {
		return false, nil
	}
	var date *rss2.RSSTime
	if len(conf.After) > 0 {
		if date, err = erss.ToRSSTime(conf.After); err != nil {
			return
		}
		if item.PubDate == nil || !item.PubDate.Time.After(date.Time) {
			return false, nil
		}
	}
	if len(conf.Before) > 0 {
		if date, err = erss.ToRSSTime(conf.Before); err != nil {
			return
		}
		if item.PubDate == nil || !item.PubDate.Time.Before(date.Time) {
			return false, nil
		}
	}
	return true, nil
}
