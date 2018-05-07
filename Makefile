.PHONEY: \
	build

build: dist/ops-log

dist/ops-log: ops-log.go
	go build -o dist/ops-log ops-log.go