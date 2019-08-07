# Usage
```console
$ rss-create --title 'My Blog' --link 'my.blog.net' \
$            --description 'My Blog about programming and cleansing products.' \
$            my_blog.rss
$ cat my_blog.rss
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
        <channel>
                <title>My Blog</title>
                <link>my.blog.net</link>
                <description>My Blog about programming and cleansing products.</description>
        </channel>
</rss>
$ rss-add-item --title "RSS's not dead!" --link 'my.blog.net/articles/1' \
$              --pubDate "$(date '+%d %b %Y %H:%M:%S %z')" my_blog.rss my_updated_blog.rss
$ cat my_updated_blog.rss
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
$ # Overwrite the original file, if you're happy with the changes:
$ mv my_updated_blog.rss my_blog.rss
$ # I forgot to specify the language of my blog, so let's change that.
$ # First I need to create a new RSS file, with the complete configuration.
$ rss-create --title 'My Blog' --link 'my.blog.net' \
$            --description 'My Blog about programming and cleansing products.' \
$            --language 'en-us' my_fixed_blog.rss
$ # Now I can copy over the items from my existing blog.
$ # Note that my_fixed_blog.rss stays unmodified; the results are
$ # written to my_fixed_and_populated_blog.rss .
$ rss-copy-items my_blog.rss my_fixed_blog.rss my_fixed_and_populated_blog.rss
$ # After making sure everything looks as expected I can overwrite the
$ # original file:
$ mv my_fixed_and_populated_blog.rss my_blog.rss && rm my_fixed_blog.rss
$ # After a while my blog grew and I want to reduce the size of my RSS
$ # file, to avoid poor performance and problems with some feed
$ # aggregators. I can simply reduce the amount of items like this:
$ rss-tail -n 32 my_blog.rss my_shrunk_blog.rss
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

To uninstall ytools call `make uninstall`.

# To think about
- `rss-sort-by-date` (for example to reorder a file after multiple calls
  of `rss-copy-items`)
- remove all the stdin stdout stuff and force the user to use files, to
  improve consistency and simplicity
