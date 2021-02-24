.PHONY: test bin/bank-ocr

default: bin/bank-ocr

test:
	go test -v ./...

bin:
	mkdir $@

bin/bank-ocr: bin
	cd cmd && \
		go build -o ../$@
