package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Tests basic functionality.
func ExampleMain1() {
	rssDummyFile := getRSSDummyFile()
	os.Args = []string{`rss-create`, `--title`, `test item`, rssDummyFile.Name()}
	main()
	printAndRemoveFile(rssDummyFile.Name())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0">
	// 	<channel>
	// 		<title>dummy</title>
	// 		<link>dummy</link>
	// 		<description>dummy</description>
	// 		<item>
	// 			<title>test item</title>
	// 		</item>
	// 	</channel>
	// </rss>
}

// (Almost) exhaustive test of the command line arguments.
func ExampleMain2() {
	rssDummyFile := getRSSDummyFile()
	os.Args = []string{`rss-create`,
		`--title`, `Venice Film Festival Tries to Quit Sinking`,
		`--link`, `http://nytimes.com/2004/12/07FEST.html`,
		`--description`, `Some of the most heated chatter at the...`,
		`--author`, `oprah@oxygen.net`,
		`--category`, `Grateful Dead`,
		`--category-domain`, ``,
		`--category`, `MSFT`,
		`--category-domain`, `http://www.fool.com/cusips`,
		`--comments`, `http://www.myblog.org/cgi-local/mt/mt-comments.cgi?entry_id=9`,
		`--enclosure-url`, `http://www.scripting.com/mp3s/weatherReportSuite.mp3`,
		`--enclosure-length`, `12216320`,
		`--enclosure-type`, `audio/mpeg`,
		`--guid`, `http://inessential.com/2002/09/01.php#a2`,
		`--guid-isPermaLink`,
		`--pubDate`, `Sun, 19 May 2002 15:21:36 GMT`,
		`--source`, `Tomalak's Realm`,
		`--source-url`, `http://www.tomalak.org/links2.xml`,
		rssDummyFile.Name()}
	main()
	printAndRemoveFile(rssDummyFile.Name())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0">
	// 	<channel>
	// 		<title>dummy</title>
	// 		<link>dummy</link>
	// 		<description>dummy</description>
	// 		<item>
	// 			<title>Venice Film Festival Tries to Quit Sinking</title>
	// 			<link>http://nytimes.com/2004/12/07FEST.html</link>
	// 			<description>Some of the most heated chatter at the...</description>
	// 			<author>oprah@oxygen.net</author>
	// 			<category>Grateful Dead</category>
	// 			<category domain="http://www.fool.com/cusips">MSFT</category>
	// 			<comments>http://www.myblog.org/cgi-local/mt/mt-comments.cgi?entry_id=9</comments>
	// 			<enclosure url="http://www.scripting.com/mp3s/weatherReportSuite.mp3" length="12216320" type="audio/mpeg"></enclosure>
	// 			<guid isPermaLink="true">http://inessential.com/2002/09/01.php#a2</guid>
	// 			<pubDate>19 May 2002 15:21:36 +0000</pubDate>
	// 			<source url="http://www.tomalak.org/links2.xml">Tomalak&#39;s Realm</source>
	// 		</item>
	// 	</channel>
	// </rss>
}

func getRSSDummyFile() *os.File {
	dummyRSS := []byte(`
		<?xml version="1.0" encoding="UTF-8"?>
		<rss version="2.0">
		        <channel>
		                <title>dummy</title>
		                <link>dummy</link>
		                <description>dummy</description>
		        </channel>
		</rss>`)
	rssDummyFile, err := ioutil.TempFile(``, `rss-add-item_test_*`)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := rssDummyFile.Write(dummyRSS); err != nil {
		log.Fatal(err)
	}
	if err := rssDummyFile.Close(); err != nil {
		log.Fatal(err)
	}
	return rssDummyFile
}

func printAndRemoveFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
	err = os.Remove(filename)
	if err != nil {
		panic(err)
	}
}
