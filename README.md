# Usage
```console
$ rss-create --title 'My Blog' --link 'my.blog.net' \
$            --description 'My Blog about programming and cleansing products.' \
$            blog.rss
$ rss-add-item --title "RSS's not dead!" --link 'my.blog.net/articles/1' \
$              --pubDate "$(date '+%d %b %Y %H:%M:%S %z')" blog.rss
$ cat blog.rss
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
        <channel>
                <title>My Blog</title>
                <link>my.blog.net</link>
                <description>My Blog about programming and cleansing products.</description>
                <item>
                        <title>RSS&#39;s not dead</title>
                        <link>my.blog.net/articles/1</link>
                        <pubDate>06 Aug 2019 18:43:28 +0200</pubDate>
                </item>
        </channel>
</rss>
$ # I forgot to specify the language of my blog, so let's change that.
$ # First I need to create a new RSS file, with the complete configuration.
$ rss-create --title 'My Blog' --link 'my.blog.net' \
$            --description 'My Blog about programming and cleansing products.' \
$            --language 'en-us' fixed_blog.rss
$ # Now I can copy over the items from my existing blog.
$ rss-copy-items blog.rss fixed_blog.rss
$ mv fixed_blog.rss blog.rss
$ # If the items of your feed have become unordered (e.g. through multiple
$ # calls of rss-copy-items), you can sort them like this:
$ rss-sort-by-date blog.rss
$ # After a while my blog grew and I want to reduce the size of my RSS
$ # file, to avoid poor performance and problems with some feed
$ # aggregators. I can simply reduce the amount of items like this:
$ rss-tail -n 32 blog.rss
```

# Installation
1. `mkdir -p "$HOME/go/src/github.com/codesoap/" && cd "$HOME/go/src/github.com/codesoap/"`
   (adapt if you've set a different `$GOPATH`)
2. `git clone https://github.com/codesoap/erss.git && cd erss/`
3. `go get ./...` to install dependencies
   ([codesoap/rss2](https://github.com/codesoap/rss2) and
   [docopt/docopt.go](https://github.com/docopt/docopt.go))
4. `make install` to install erss (if you just want the binaries do
   `make all`)

To uninstall erss call `make uninstall`.
