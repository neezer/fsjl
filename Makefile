.PHONEY: build clean

build:  dist/ops-log_darwin_amd64 \
	dist/ops-log_linux_amd64 \
	dist/ops-log_freebsd_amd64 \
	dist/ops-log_openbsd_amd64

dist/ops-log_darwin_amd64 dist/ops-log_linux_amd64 ops-log_freebsd_amd64 dist/ops-log_openbsd_amd64: ops-log.go
	go get github.com/mitchellh/gox
	gox -osarch="darwin/amd64 linux/amd64 freebsd/amd64 openbsd/amd64" -output="dist/ops-log_{{.OS}}_{{.Arch}}"

clean:
	rm -rf dist
