.PHONY: all install uninstall clean

PREFIX = /usr/local
MANPREFIX = /usr/local/man
ROOT = github.com/codesoap/erss

all: bin/rss-create bin/rss-add-item bin/rss-copy-items bin/rss-sort-by-date \
     bin/rss-tail

install: all
	mkdir -p                              "${DESTDIR}${PREFIX}/bin"
	install -m 755 "bin/rss-create"       "${DESTDIR}${PREFIX}/bin"
	install -m 755 "bin/rss-add-item"     "${DESTDIR}${PREFIX}/bin"
	install -m 755 "bin/rss-copy-items"   "${DESTDIR}${PREFIX}/bin"
	install -m 755 "bin/rss-sort-by-date" "${DESTDIR}${PREFIX}/bin"
	install -m 755 "bin/rss-tail"         "${DESTDIR}${PREFIX}/bin"

uninstall:
	rm -f "${DESTDIR}${PREFIX}/bin/rss-create" \
	      "${DESTDIR}${PREFIX}/bin/rss-add-item" \
	      "${DESTDIR}${PREFIX}/bin/rss-copy-items" \
	      "${DESTDIR}${PREFIX}/bin/rss-sort-by-date" \
	      "${DESTDIR}${PREFIX}/bin/rss-tail" \

clean:
	rm -rf bin

bin/rss-create: cmd/rss-create/main.go lib.go
	go build -o bin/rss-create ${ROOT}/cmd/rss-create

bin/rss-add-item: cmd/rss-add-item/main.go lib.go
	go build -o bin/rss-add-item ${ROOT}/cmd/rss-add-item

bin/rss-copy-items: cmd/rss-copy-items/main.go lib.go
	go build -o bin/rss-copy-items ${ROOT}/cmd/rss-copy-items

bin/rss-sort-by-date: cmd/rss-sort-by-date/main.go lib.go
	go build -o bin/rss-sort-by-date ${ROOT}/cmd/rss-sort-by-date

bin/rss-tail: cmd/rss-tail/main.go lib.go
	go build -o bin/rss-tail ${ROOT}/cmd/rss-tail
