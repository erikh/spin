test: gen
	go test -v ./...

clean: 
	rm -f spin-broker.db

stop-server:
	pkill spin-broker || :
	sleep 1

server: stop-server
	GOBIN=${PWD}/bin go install -v ./...
	bin/spin-broker &

gen:
	go run goa.design/goa/v3/cmd/goa gen code.hollensbe.org/erikh/spin/design

examples:
	go run goa.design/goa/v3/cmd/goa example code.hollensbe.org/erikh/spin/design

.PHONY: gen
