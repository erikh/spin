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
`/tmp/host-path-test`

```bash
#!bash

go run ./cmd/agents/sa-host-path &

pkg=$(go run ./cmd/spin-cli spin-broker new | jq -r .)

go run ./cmd/spin-cli spin-broker add --id "${pkg}" --body "{
  \"action\": \"add_volume\",
  \"parameters\": { \"path\": \"test\" },
  \"resource\": \"storage\"
}"

go run ./cmd/spin-cli spin-broker add --id "${pkg}" --body "{
  \"action\": \"create_image\",
  \"parameters\": { \"volume_path\": \"test\", \"image_name\": \"test.img\", \"image_size\": \"50\" },
  \"resource\": \"storage\"
}"

go run ./cmd/spin-cli spin-broker enqueue --id "${pkg}"
sleep 1
pkill sa-host-path
```
