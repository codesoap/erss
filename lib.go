package erss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codesoap/rss2"
)

func ToRSSTime(dateString string) (t *rss2.RSSTime, err error) {
	date, err := rss2.ParseRSSTime(dateString)
	return &date, err
}

func ToCategory(category, categoryDomain string) (c *rss2.Category, err error) {
	if c, err = rss2.NewCategory(category); err != nil {
		return
	}
	c.Domain = categoryDomain
	return
}

func WriteRSS(rss *rss2.RSS, outfile_name string) (err error) {
	rss_bytes, err := xml.MarshalIndent(rss, ``, "\t")
	if err != nil {
		return fmt.Errorf(`error when rendering rss: %s`, err.Error())
	}
	var outfile *os.File
	outfile, err = os.Create(outfile_name)
	if err != nil {
		return
	}
	defer outfile.Close()
	fmt.Fprint(outfile, xml.Header)
	fmt.Fprintln(outfile, string(rss_bytes))
	return
}

func ReadRSS(filename string) (r *rss2.RSS, err error) {
	var input []byte
	if input, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	var rss rss2.RSS
	err = xml.Unmarshal(input, &rss)
	return &rss, err
}
