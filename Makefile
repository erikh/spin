test: gen
	go test -v ./...

autotest:
	@echo 1>&2
	@echo 1>&2 '`go get github.com/cespare/reflex` to use this task'
	@echo 1>&2
	reflex -r '\.go$$' -R '^(gen|goa[0-9]+)/' make test

test-all: gen
	go test -v ./... -count 1

clean: 
	rm -f spin-broker.db

stop-server:
	pkill spin-broker || :
	sleep 1

server: stop-server
	GOBIN=${PWD}/bin go install -v ./...
	bin/spin-broker start &

gen:
	go run goa.design/goa/v3/cmd/goa gen code.hollensbe.org/erikh/spin/design

examples:
	go run goa.design/goa/v3/cmd/goa example code.hollensbe.org/erikh/spin/design

.PHONY: gen
