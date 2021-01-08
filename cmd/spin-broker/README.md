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

## Idempotency in a tight package

Queues are modeled through `bbolt` buckets which uses operations that are
isolated in transactions. Each queue is assigned to a resource - a string that
identifies the kind of work they are supposed to be performing. A queue is a
collection of "commands".

A command contains a resource, action, parameters, and dependencies:

- resources are mentioned above; they exist to separate powers between agents,
  which will be discussed more below.
- actions are the command name to execute; `argv[0]` if you will.
- parameters are a key/value map of parameters to send along with the action to
  the agent. These are typed, and validated for presence as well.
- dependencies are a list of identifiers (UUIDs currently) that must be
  satisfied before this queue item can be yielded.

Agents are units of code which read from the broker and carry out its actions,
one at a time (currently, at least!). Agents consume from a single queue, and
model an actual resource (storage, network, compute) in behavior.

For example, the storage agent has a `provision volume` action and a
`create image` action, as well as removal actions. A host path agent implements
this, but there is no reason a zfs agent cannot be built. Agents do not
communicate back to the broker, they simply do their work and move on.

The [API server](../spin-apiserver) communicates with the broker to issue
commands to the agents. If you just want to launch a VM, talk to the API
server.

To issue entries into the queue, you must organize them into a "package". The
package contains your mini-queue and must be enqueued as a group to enter
processing, otherwise it dies. Add commands to your package and then enqueue
the package. After enqueuing, you can poll the status for a result. A result of
`record_not_found` (type `*goa.ServiceError` in the design) will indicate that
the queue is simply empty, and then you can re-poll until something comes.
Check out the agent `Loop()` implementation for more.

## Toying with the agents

`make server` then run this script; the files/directories will appear in
`~/.config/spin` and can be safely deleted after this is done running.

This will create the storage and the vm supervisory files, and start the vm.
Then, it will shut it down first gracefully, then forcefully after 5 seconds.

Please note that the JSON in the example is in an extreme state of flux and may
become out of date rapidly. Check [the emulation agent's source](../../pkg/resources/emulation)
for the most updated information on the subject.

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
  "Images": [
    {
      "Path": "/home/erikh/.config/spin/images/test/test.raw",
      "Cdrom": false,
      "Volume": "test"
    },
    {
      "Volume": "test",
      "Path": "/home/erikh/vm-images/isos/manjaro-architect-20.0.3-200607-linux56.iso",
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
