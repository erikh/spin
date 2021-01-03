# spin-broker: a messaging system for FSMs

spin-broker is a message broker that has the following guarantees:

- serial message delivery, in order, for all tasks with no outstanding dependencies
- pipelined queue staging (group execution)
- resource named queues, for interacting with agents
- dependencies system, allowing you to make tasks dependent on each other
  without impacting queue delivery for head-of-line
- group and task-level status querying
- reliable message delivery
- pretty fast

spin-broker has the following properties:

- packages are collections of commands that are enqueued as a group.
  - enqueuing the package is the act of distributing it to the resource queues.
- commands are a collection of `{ resource, action, parameters, dependent commands }`.
  - They yield statuses, which are true/false values based on the success of
    the execution of the action. Parameters are used to variate the action.
- queues are lists of commands, keyed by resource. They are returned in the
  order they are enqueued, as long as their dependencies are resolved.
- agents are expected to listed on queues named after the resource they manage.
  An example is in [sa-host-path](../agents/sa-host-path).

## Toying with the host-path agent

`make server` then run this script; the files/directories will appear in
`~/.config/spin` and can be safely deleted after this is done running.

```bash
#!env bash

set -xe

pkill sa-host-path || :
pkill sa-emulation || :

go run ./cmd/agents/sa-host-path &
go run ./cmd/agents/sa-emulation &

pkg=$(go run ./cmd/spin-broker message new)

go run ./cmd/spin-broker message add "${pkg}" storage add_volume path=test
go run ./cmd/spin-broker message add "${pkg}" storage create_image volume_path=test image_name=test.raw image_size=50
go run ./cmd/spin-broker message add "${pkg}" emulation write_config id=1 'vm={
  "Name": "vm",
  "Cpus": 8,
  "Memory": 8192,
  "Storage": [
    {
      "Volume": "test",
      "Image": "/home/erikh/.config/spin/images/test/test.raw",
      "ImageSize": 50
    },
    {
      "Volume": "test",
      "Image": "/home/erikh/vm-images/isos/manjaro-architect-20.0.3-200607-linux56.iso",
      "ImageSize": 50,
      "Cdrom": true
    }
  ]
}'
go run ./cmd/spin-broker message add "${pkg}" emulation start id=1
go run ./cmd/spin-broker message enqueue "${pkg}"

sleep 15

go run ./cmd/spin-broker message status "${pkg}"

pkg=$(go run ./cmd/spin-broker message new)
go run ./cmd/spin-broker message add "${pkg}" emulation stop id=1
go run ./cmd/spin-broker message enqueue "${pkg}"

sleep 1

pkill sa-host-path || :
pkill sa-emulation || :

go run ./cmd/spin-broker message status "${pkg}"
```
