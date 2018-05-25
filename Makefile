.PHONEY: build clean

build:  dist/fsjl_darwin_amd64 \
	dist/fsjl_linux_amd64 \
	dist/fsjl_freebsd_amd64 \
	dist/fsjl_openbsd_amd64

dist/fsjl_darwin_amd64 dist/fsjl_linux_amd64 fsjl_freebsd_amd64 dist/fsjl_openbsd_amd64: fsjl.go
	go get github.com/mitchellh/gox
	gox -osarch="darwin/amd64 linux/amd64 freebsd/amd64 openbsd/amd64" -output="dist/fsjl_{{.OS}}_{{.Arch}}"

clean:
	rm -rf dist
