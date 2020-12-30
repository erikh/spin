GOLANGCI_LINT_VERSION := 1.34.1

test: gen lint
	go test -v ./...

autotest:
	@echo 1>&2
	@echo 1>&2 '`make update-reflex` to use this task'
	@echo 1>&2
	bin/reflex -r '\.go$$' -R '^(gen|goa[0-9]+)/' make test

test-all: gen
	go test -v ./... -count 1

clean: 
	rm -f spin-broker.db spin-registry.db

stop-server:
	pkill spin-broker || :
	pkill spin-registry || :
	sleep 1

install-local:
	GOBIN=${PWD}/bin go install -v ./...

server: stop-server install-local 
	bin/spin-broker start &
	bin/spin-registry start &

gen:
	go run goa.design/goa/v3/cmd/goa gen code.hollensbe.org/erikh/spin/design

examples:
	go run goa.design/goa/v3/cmd/goa example code.hollensbe.org/erikh/spin/design

lint:
	bin/golangci-lint run -v

install-deps: update-golangci-lint update-reflex

update-golangci-lint:
	wget -O- https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64.tar.gz | tar vxz --strip-components=1 -C bin golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64/golangci-lint

update-reflex:
	GO111MODULE=off GOBIN=${PWD}/bin go get -u github.com/cespare/reflex

.PHONY: gen
