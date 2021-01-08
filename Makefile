GOLANGCI_LINT_VERSION := 1.34.1

test: gen statik-build lint
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

stop-servers:
	pkill spin-apiserver || :
	pkill sa- || :
	pkill spin-registry || :
	pkill spin-broker || :
	sleep 1

react-install:
	cd ui && yarn install

react-build:
	cd ui && yarn build

statik-build: react-build
	bin/statik -src=ui/build -f

install: test
	go install -v ./...

install-local: statik-build
	GOBIN=${PWD}/bin go install -v ./...

servers: stop-servers install-local
	bin/spin-broker start &
	bin/spin-registry start &
	bin/spin-apiserver start &
	bin/sa-emulation &
	bin/sa-host-path &

gen:
	go run goa.design/goa/v3/cmd/goa gen github.com/erikh/spin/design
	goimports -w $$(find . -name '*.go' -type f)

examples:
	go run goa.design/goa/v3/cmd/goa example github.com/erikh/spin/design

lint:
	bin/golangci-lint run -v

install-deps: update-golangci-lint update-reflex update-statik react-install

update-golangci-lint:
	wget -O- https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64.tar.gz | tar vxz --strip-components=1 -C bin golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64/golangci-lint

update-reflex:
	GO111MODULE=off GOBIN=${PWD}/bin go get -u github.com/cespare/reflex

update-statik:
	GO111MODULE=off GOBIN=${PWD}/bin go get -u github.com/rakyll/statik

.PHONY: gen
