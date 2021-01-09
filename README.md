# spin: a framework/tools for managing VM configurations

spin allows you to manage VMs on your linux bare metal host with ease. Powered
by [qemu](https://www.qemu.org) and
[systemd](https://www.freedesktop.org/wiki/Software/systemd/), it stands upon
reliable technology providing its own framework for applying configuration in a
safely concurrent, idempotent way, allowing you to perform many small
operations on VMs in a quick and efficient way.

In short, we do the high level stuff and leave the big iron tools doing what
they do best.

spin is fully backed by a [openapi specification](gen/http/openapi3.yaml)
for its API calls, from the high level to the blow-your-foot-off calls.
Generate your own, or check out our [clients](clients) directory for a good
set of golang ones.

## Installation and Usage

For installations, we only support "development mode" at the moment. For a
quick start, try this script on for size. You will need the following tools to
build the binaries:

- [golang](golang.org)
- [yarn](https://yarnpkg.com/)

Most other things will be installed by the `make install-deps` task automatically.

```bash
#!/usr/bin/env bash

go get -d github.com/erikh/spin
cd $GOPATH/github.com/erikh/spin
make install-deps
make servers
make install-local
```

Binaries will be in `bin/`.

To interact with the `spin` tool (which has help; try `spin help`), try the
following ideas:

(please note the servers must be running, provided by `make servers`)

```bash
# these first few steps are necessary to create vms for now, soon this will not be
mkdir -p ~/.config/spin/images/test

# first argument is the name of the volume, second is the path to the volume's files
spin-registry message storage volume create test ~/.config/spin/images/test

# the second argument (`test`) here is the volume name provided above. the
# machine's image will be put in this directory under the name `my_cool_image`.
spin vm create --cdrom ~/path/to/some.iso my_cool_image test
1

# the id above is the id of the vm newly created; it can be started like so:
spin start 1

# to access it, you can get a console by starting the console server and
# viewing a specific console (will open a link in your browser):
spin serve &
spin view 1

# stop and shutdown have a distinction; stop will force terminate through
# systemd if shutdown doesn't succeed; shutdown will merely try to stop it
# through normal ACPI shutdown signals.
spin shutdown 1
spin stop 1

##
# some other commands:
##
spin vm list              # lists all vms by id
spin vm image list 1      # lists attached devices to my_cool_image
spin vm image detach 1 1  # detaches cdrom, index 1 on my_cool_image
```

Run `make stop-servers` to stop the servers when you are done playing! :D

### Development

Just some notes:

goa gen is done on each `make test` run, that includes the autotester, so you
can use it to refactor designs fairly quick.

golangci-lint is required to run the tests. `make install-deps` to install it
and reflex.

`make autotest` will start an autotester using
[reflex](https://github.com/cespare/reflex). It's very handy!

The tests are currently pretty safe to run, they don't generate in any place
other than `/tmp` and the project directory. `.gitignore` is programmed
properly, too. `test` and autotest leverage the golang test cache for faster
response time, but you can `make test-all` to force a run of all tests.

## Architecture

The architecture is a mix of the following technologies:

- [golang](https://golang.org) is ... you know what golang is.
- [goa](https://goa.design) is a framework for designing REST and GRPC APIs.
- [bbolt](https://pkg.go.dev/go.etcd.io/bbolt/) is a key/value database system similar to LMDB, just for Golang.

And it contains the following sub-components:

- [spin-broker](cmd/spin-broker) arranges messages so that they are processed
  in queues monitored by agents. These agents carry out the work of the
  instructions carried in the queue.
- [spin-registry](cmd/spin-registry) holds the values of entities. Soon, this
  will include journalling the changes of entities.
- [spin-apiserver](cmd/spin-apiserver) is the front API (the "brains") that
  orchestrates the interaction between the registry, broker and the user.

## Author

Erik Hollensbe <github@hollensbe.org>
