package erss

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/codesoap/rss2"
)

func ToRSSTime(dateString string) (t *rss2.RSSTime, err error) {
	date, err := rss2.ParseRSSTime(dateString)
	return &date, err
}

func PrintOrWriteResult(rss *rss2.RSS, outfile_name string) (err error) {
	rss_bytes, err := xml.MarshalIndent(rss, ``, "\t")
	if err != nil {
		return fmt.Errorf(`error when rendering rss: %s`, err.Error())
	}
	if len(outfile_name) == 0 {
		fmt.Print(xml.Header)
		fmt.Println(string(rss_bytes))
	} else {
		var outfile *os.File
		outfile, err = os.Create(outfile_name)
		if err != nil {
			return
		}
		defer outfile.Close()
		fmt.Fprint(outfile, xml.Header)
		fmt.Fprintln(outfile, string(rss_bytes))
	}
	return
}
