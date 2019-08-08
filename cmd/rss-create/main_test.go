package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Tests basic functionality.
func ExampleMain1() {
	os.Args = []string{`rss-create`,
		`--title`, `test title`,
		`--link`, `link.com`,
		`--description`, `special characters: " & ;`,
		`main_test.rss`}
	main()
	printAndRemoveFile(`main_test.rss`)
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0">
	// 	<channel>
	// 		<title>test title</title>
	// 		<link>link.com</link>
	// 		<description>special characters: &#34; &amp; ;</description>
	// 	</channel>
	// </rss>
}

// (Almost) exhaustive test of the command line arguments.
func ExampleMain2() {
	os.Args = []string{`rss-create`,
		`--title`, `GoUpstate.com News Headlines`,
		`--link`, `http://www.goupstate.com/`,
		`--description`, `The latest news from GoUpstate.com, a Spartanburg Herald-Journal Web site.`,
		`--language`, `en-us`,
		`--copyright`, `Copyright 2002, Spartanburg Herald-Journal`,
		`--managingEditor`, `geo@herald.com (George Matesky)`,
		`--webMaster`, `betty@herald.com (Betty Guernsey`,
		`--pubDate`, `01 Jul 1999 12:15 B`,
		`--lastBuildDate`, `Sat, 07 Sep 2002 00:00:01 GMT`,
		`--category`, `Grateful Dead`,
		`--category-domain`, ``,
		`--category`, `MSFT`,
		`--category-domain`, `http://www.fool.com/cusips`,
		`--generator`, `MightyInHouse Content System v2.3`,
		`--docs`, `http://blogs.law.harvard.edu/tech/rss`,
		`--cloud-domain`, `rpc.sys.com`,
		`--cloud-port`, `80`,
		`--cloud-path`, `/RPC2`,
		`--cloud-registerProcedure`, `myCloud.rssPleaseNotify`,
		`--cloud-protocol`, `xml-rpc`,
		`--ttl`, `120`,
		`--image-url`, `my.im.ag/e.png`,
		`--image-title`, `GoUpstate.com News Headlines`,
		`--image-link`, `http://www.goupstate.com/`,
		`--image-width`, `60`,
		`--image-height`, `80`,
		`--image-description`, `my image description`,
		`--rating`, `my rating`,
		`--textInput-title`, `my textInput title`,
		`--textInput-description`, `my textInput description`,
		`--textInput-name`, `my textInput name`,
		`--textInput-link`, `my.textInput.li/nk.cgi`,
		`--skipHour`, `0`,
		`--skipHour`, `4`,
		`--skipDay`, `Sunday`,
		`--skipDay`, `Saturday`,
		`main_test.rss`}
	main()
	printAndRemoveFile(`main_test.rss`)
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rss version="2.0">
	// 	<channel>
	// 		<title>GoUpstate.com News Headlines</title>
	// 		<link>http://www.goupstate.com/</link>
	// 		<description>The latest news from GoUpstate.com, a Spartanburg Herald-Journal Web site.</description>
	// 		<language>en-us</language>
	// 		<copyright>Copyright 2002, Spartanburg Herald-Journal</copyright>
	// 		<managingEditor>geo@herald.com (George Matesky)</managingEditor>
	// 		<webMaster>betty@herald.com (Betty Guernsey</webMaster>
	// 		<pubDate>01 Jul 1999 12:15:00 +0200</pubDate>
	// 		<lastBuildDate>07 Sep 2002 00:00:01 +0000</lastBuildDate>
	// 		<category>Grateful Dead</category>
	// 		<category domain="http://www.fool.com/cusips">MSFT</category>
	// 		<generator>MightyInHouse Content System v2.3</generator>
	// 		<docs>http://blogs.law.harvard.edu/tech/rss</docs>
	// 		<cloud domain="rpc.sys.com" port="80" path="/RPC2" registerProcedure="myCloud.rssPleaseNotify" protocol="xml-rpc"></cloud>
	// 		<ttl>120</ttl>
	// 		<image>
	// 			<url>my.im.ag/e.png</url>
	// 			<title>GoUpstate.com News Headlines</title>
	// 			<link>http://www.goupstate.com/</link>
	// 			<width>60</width>
	// 			<height>80</height>
	// 			<description>my image description</description>
	// 		</image>
	// 		<rating>my rating</rating>
	// 		<textInput>
	// 			<title>my textInput title</title>
	// 			<description>my textInput description</description>
	// 			<name>my textInput name</name>
	// 			<link>my.textInput.li/nk.cgi</link>
	// 		</textInput>
	// 		<skipHours>
	// 			<hour>0</hour>
	// 			<hour>4</hour>
	// 		</skipHours>
	// 		<skipDays>
	// 			<day>Sunday</day>
	// 			<day>Saturday</day>
	// 		</skipDays>
	// 	</channel>
	// </rss>
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
