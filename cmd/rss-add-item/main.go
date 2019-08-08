package main

import (
	"fmt"
	"os"

	"github.com/codesoap/erss"
	"github.com/codesoap/rss2"
	"github.com/docopt/docopt-go"
)

// --category-domain is not optional, when --category is used, to avoid
// problems when --category is given multiple times. An empty string can
// be passed to --category-domain, if it's not needed.
var usage = `
Add an RSS item to an existing RSS 2.0 file.
Note that at least --title or --description must be given.

Usage:
    rss-add-item [--title=<title>]
                 [--link=<link>]
                 [--description=<desc>]
                 [--author=<author>]
                 [(  --category=<category>
                     --category-domain=<domain>
                 )...]
                 [--comments=<comments>]
                 [(  --enclosure-url=<url>
                     --enclosure-length=<length>
                     --enclosure-type=<type>
                 )]
                 [(  --guid=<guid>
                     [--guid-isPermaLink]
                 )]
                 [--pubDate=<date>]
                 [(  --source=<source>
                     --source-url=<url>
                 )]
                 <file>
`

type conf struct {
	Title           string
	Link            string
	Description     string
	Author          string
	Category        []string
	CategoryDomain  []string
	Comments        string
	EnclosureUrl    string
	EnclosureLength int
	EnclosureType   string
	Guid            string
	GuidIsPermaLink bool
	PubDate         string
	Source          string
	SourceUrl       string
	File            string
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
	if len(conf.Title) == 0 && len(conf.Description) == 0 {
		fmt.Fprintln(os.Stderr, `At least --title or --description must be provided.`)
		os.Exit(2)
	}
	rss, err := erss.ReadRSS(conf.File)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when reading the existing RSS:`, err.Error())
		os.Exit(3)
	}
	item, err := getItem(&conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when creating item:`, err.Error())
		os.Exit(4)
	}
	rss.Channel.Items = append(rss.Channel.Items, item)
	if err = erss.WriteResult(rss, conf.File); err != nil {
		fmt.Fprintln(os.Stderr, `Error when rendering:`, err.Error())
		os.Exit(5)
	}
}

func getItem(conf *conf) (item *rss2.Item, err error) {
	if item, err = rss2.NewItem(conf.Title, conf.Description); err != nil {
		return
	}
	item.Link = conf.Link
	item.Author = conf.Author
	for i := 0; i < len(conf.Category); i++ {
		var category *rss2.Category
		category, err = erss.ToCategory(conf.Category[i], conf.CategoryDomain[i])
		if err != nil {
			return
		} else if i == 0 {
			item.Categories = []*rss2.Category{category}
		} else {
			item.Categories = append(item.Categories, category)
		}
	}
	item.Comments = conf.Comments
	if err = addEnclosure(item, conf); err != nil {
		return
	}
	if err = addGUID(item, conf); err != nil {
		return
	}
	if len(conf.PubDate) > 0 {
		if item.PubDate, err = erss.ToRSSTime(conf.PubDate); err != nil {
			return
		}
	}
	err = addSource(item, conf)
	return
}

func addEnclosure(item *rss2.Item, conf *conf) (err error) {
	if len(conf.EnclosureUrl) > 0 {
		item.Enclosure, err = rss2.NewEnclosure(conf.EnclosureUrl,
			conf.EnclosureLength, conf.EnclosureType)
	}
	return
}

func addGUID(item *rss2.Item, conf *conf) (err error) {
	if len(conf.Guid) > 0 {
		if item.GUID, err = rss2.NewGUID(conf.Guid); err != nil {
			return
		}
		item.GUID.IsPermaLink = conf.GuidIsPermaLink
	}
	return
}

func addSource(item *rss2.Item, conf *conf) (err error) {
	if len(conf.Source) > 0 {
		item.Source, err = rss2.NewSource(conf.Source, conf.SourceUrl)
	}
	return
}
