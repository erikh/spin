# API Clients

This goes over the various imports you can use to interact with the services
spin provides. It is generally not recommended to use the rest of the tree
outside of `clients/`.

- [api](api) is for the [spin-apiserver](../cmd/spin-apiserver), and is probably what you want to use.
- [broker](broker) imports as `brokerclient` and is for the spin-broker, which
  is described [here](../cmd/spin-broker).
- [registry](registry) imports as `registryclient` and is for the [spin-registry](../cmd/spin-registry).
