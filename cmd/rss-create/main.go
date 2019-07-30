package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/codesoap/erss"
	"github.com/codesoap/rss2"
	"github.com/docopt/docopt-go"
)

var usage = `
Create an RSS 2.0 file. The file will contain no RSS items.

Usage:
    rss-create --title=<title>
               --link=<link>
               --description=<desc>
               [--language=<lang>]
               [--copyright=<notice>]
               [--managingEditor=<name>]
               [--webMaster=<name>]
               [--pubDate=<date>]
               [--lastBuildDate=<date>]
               [(  --category=<category>
                   [--category-domain=<domain>]
               )]
               [--generator=<generator>]
               [--docs=<docs>]
               [(  --cloud-domain=<domain>
                   --cloud-port=<port>
                   --cloud-path=<path>
                   --cloud-registerProcedure=<rp>
                   --cloud-protocol=<protocol>
               )]
               [--ttl=<ttl>]
               [(  --image-url=<url>
                   --image-title=<title>
                   --image-link=<link>
                   [--image-width=<width>]
                   [--image-height=<height>]
                   [--image-description=<desc>]
               )]
               [--rating=<rating>]
               [(  --textInput-title=<title>
                   --textInput-description=<desc>
                   --textInput-name=<name>
                   --textInput-link=<link>
               )]
               [(--skipHour=<hour>)...]
               [(--skipDay=<day>)...]
               [<outfile>]
`

type conf struct {
	Title                  string
	Link                   string
	Description            string
	Language               string
	Copyright              string
	ManagingEditor         string
	WebMaster              string
	PubDate                string
	LastBuildDate          string
	Category               string
	CategoryDomain         string
	Generator              string
	Docs                   string
	CloudDomain            string
	CloudPort              int
	CloudPath              string
	CloudRegisterProcedure string
	CloudProtocol          string
	Ttl                    int
	ImageUrl               string
	ImageTitle             string
	ImageLink              string
	ImageWidth             int
	ImageHeight            int
	ImageDescription       string
	Rating                 string
	TextInputTitle         string
	TextInputDescription   string
	TextInputName          string
	TextInputLink          string
	// SkipHour should be []int, but docopt-go doesn't support it.
	SkipHour []string
	SkipDay  []string
	Outfile  string
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
	channel, err := rss2.NewChannel(conf.Title, conf.Link, conf.Description)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when creating RSS channel:`, err.Error())
		os.Exit(3)
	}
	err = addOptionalElements(channel, &conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Error when parsing arguments:`, err.Error())
		os.Exit(3)
	}
	rss := rss2.NewRSS(channel)
	if err = erss.PrintOrWriteResult(rss, conf.Outfile); err != nil {
		fmt.Fprintln(os.Stderr, `Error when rendering:`, err.Error())
		os.Exit(4)
	}
}

func addOptionalElements(channel *rss2.Channel, conf *conf) (err error) {
	channel.Language = conf.Language
	channel.Copyright = conf.Copyright
	channel.ManagingEditor = conf.ManagingEditor
	channel.WebMaster = conf.WebMaster
	if len(conf.PubDate) > 0 {
		if channel.PubDate, err = erss.ToRSSTime(conf.PubDate); err != nil {
			return
		}
	}
	if len(conf.LastBuildDate) > 0 {
		channel.LastBuildDate, err = erss.ToRSSTime(conf.LastBuildDate)
		if err != nil {
			return
		}
	}
	if len(conf.Category) > 0 {
		channel.Category, err = erss.ToCategory(conf.Category, conf.CategoryDomain)
		if err != nil {
			return
		}
	}
	channel.Generator = conf.Generator
	channel.Docs = conf.Docs
	if err = addCloud(channel, conf); err != nil {
		return
	}
	channel.TTL = conf.Ttl
	if err = addImage(channel, conf); err != nil {
		return
	}
	channel.Rating = conf.Rating
	if err = addTextInput(channel, conf); err != nil {
		return
	}
	if err = addSkipHours(channel, conf); err != nil {
		return
	}
	err = addSkipDays(channel, conf)
	return
}

func addCloud(channel *rss2.Channel, conf *conf) (err error) {
	if len(conf.CloudDomain) > 0 {
		channel.Cloud, err = rss2.NewCloud(conf.CloudDomain, conf.CloudPort,
			conf.CloudPath, conf.CloudRegisterProcedure, conf.CloudProtocol)
	}
	return
}

func addTextInput(channel *rss2.Channel, conf *conf) (err error) {
	if len(conf.TextInputTitle) > 0 {
		channel.TextInput, err = rss2.NewTextInput(conf.TextInputTitle,
			conf.TextInputDescription, conf.TextInputName, conf.TextInputLink)
	}
	return
}

func addImage(channel *rss2.Channel, conf *conf) (err error) {
	if len(conf.ImageUrl) > 0 {
		channel.Image, err = rss2.NewImage(conf.ImageUrl, conf.ImageTitle,
			conf.ImageLink)
		if err != nil {
			return
		}
		if conf.ImageWidth > 0 {
			channel.Image.Width = conf.ImageWidth
		}
		if conf.ImageHeight > 0 {
			channel.Image.Height = conf.ImageHeight
		}
		if len(conf.ImageDescription) > 0 {
			channel.Image.Description = conf.ImageDescription
		}
	}
	return
}

func addSkipHours(channel *rss2.Channel, conf *conf) (err error) {
	if len(conf.SkipHour) > 0 {
		var hours []int
		for _, hour := range conf.SkipHour {
			var hourInt int
			hourInt, err = strconv.Atoi(hour)
			if err != nil {
				return
			}
			hours = append(hours, hourInt)
		}
		var skipHours *rss2.SkipHours
		if skipHours, err = rss2.NewSkipHours(hours); err != nil {
			return
		}
		channel.SkipHours = skipHours
	}
	return
}

func addSkipDays(channel *rss2.Channel, conf *conf) (err error) {
	if len(conf.SkipDay) > 0 {
		var days []time.Weekday
		daysLookup := map[string]time.Weekday{
			`Sunday`:    time.Sunday,
			`Monday`:    time.Monday,
			`Tuesday`:   time.Tuesday,
			`Wednesday`: time.Wednesday,
			`Thursday`:  time.Thursday,
			`Friday`:    time.Friday,
			`Saturday`:  time.Saturday,
		}
		for _, day := range conf.SkipDay {
			var dayWeekday time.Weekday
			dayWeekday, ok := daysLookup[day]
			if !ok {
				return fmt.Errorf(`invalid weekday: %s`, day)
			}
			days = append(days, dayWeekday)
		}
		var skipDays *rss2.SkipDays
		skipDays = rss2.NewSkipDays(days)
		channel.SkipDays = skipDays
	}
	return
}
