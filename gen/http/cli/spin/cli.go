// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin HTTP client CLI support package
//
// Command:
// $ goa gen github.com/erikh/spin/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	spinapiserverc "github.com/erikh/spin/gen/http/spin_apiserver/client"
	spinbrokerc "github.com/erikh/spin/gen/http/spin_broker/client"
	spinregistryc "github.com/erikh/spin/gen/http/spin_registry/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `spin-broker (new|add|enqueue|status|next|complete)
spin-apiserver (vm-create|vm-delete|vm-list|vm-get|vm-update|control-start|control-stop|control-shutdown)
spin-registry (vm-create|vm-update|vm-delete|vm-get|vm-list|storage-volumes-list|storage-volumes-create|storage-volumes-delete|storage-images-list|storage-images-create|storage-images-delete|storage-images-get)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` spin-broker new` + "\n" +
		os.Args[0] + ` spin-apiserver vm-create --body '{
      "cpus": 17092856392700368829,
      "memory": 497222928895641055,
      "name": "Aspernatur vel vel illum voluptatem voluptatibus est.",
      "storage": [
         {
            "cdrom": true,
            "image": "Esse reprehenderit qui molestias eum voluptatem.",
            "image_size": 11533824901793082147,
            "volume": "Occaecati deserunt qui praesentium."
         },
         {
            "cdrom": true,
            "image": "Esse reprehenderit qui molestias eum voluptatem.",
            "image_size": 11533824901793082147,
            "volume": "Occaecati deserunt qui praesentium."
         }
      ]
   }'` + "\n" +
		os.Args[0] + ` spin-registry vm-create --body '{
      "cpus": 2551606264180670795,
      "images": [
         {
            "cdrom": true,
            "path": "Dignissimos qui error modi.",
            "volume": "Corrupti et voluptatibus et et occaecati."
         },
         {
            "cdrom": true,
            "path": "Dignissimos qui error modi.",
            "volume": "Corrupti et voluptatibus et et occaecati."
         },
         {
            "cdrom": true,
            "path": "Dignissimos qui error modi.",
            "volume": "Corrupti et voluptatibus et et occaecati."
         },
         {
            "cdrom": true,
            "path": "Dignissimos qui error modi.",
            "volume": "Corrupti et voluptatibus et et occaecati."
         }
      ],
      "memory": 17694085759044319548,
      "name": "Nobis quia."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		spinBrokerFlags = flag.NewFlagSet("spin-broker", flag.ContinueOnError)

		spinBrokerNewFlags = flag.NewFlagSet("new", flag.ExitOnError)

		spinBrokerAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		spinBrokerAddBodyFlag = spinBrokerAddFlags.String("body", "REQUIRED", "")
		spinBrokerAddIDFlag   = spinBrokerAddFlags.String("id", "REQUIRED", "Package ID")

		spinBrokerEnqueueFlags  = flag.NewFlagSet("enqueue", flag.ExitOnError)
		spinBrokerEnqueueIDFlag = spinBrokerEnqueueFlags.String("id", "REQUIRED", "Package ID")

		spinBrokerStatusFlags  = flag.NewFlagSet("status", flag.ExitOnError)
		spinBrokerStatusIDFlag = spinBrokerStatusFlags.String("id", "REQUIRED", "Package ID")

		spinBrokerNextFlags        = flag.NewFlagSet("next", flag.ExitOnError)
		spinBrokerNextResourceFlag = spinBrokerNextFlags.String("resource", "REQUIRED", "resource type")

		spinBrokerCompleteFlags    = flag.NewFlagSet("complete", flag.ExitOnError)
		spinBrokerCompleteBodyFlag = spinBrokerCompleteFlags.String("body", "REQUIRED", "")

		spinApiserverFlags = flag.NewFlagSet("spin-apiserver", flag.ContinueOnError)

		spinApiserverVMCreateFlags    = flag.NewFlagSet("vm-create", flag.ExitOnError)
		spinApiserverVMCreateBodyFlag = spinApiserverVMCreateFlags.String("body", "REQUIRED", "")

		spinApiserverVMDeleteFlags  = flag.NewFlagSet("vm-delete", flag.ExitOnError)
		spinApiserverVMDeleteIDFlag = spinApiserverVMDeleteFlags.String("id", "REQUIRED", "ID of VM to delete")

		spinApiserverVMListFlags = flag.NewFlagSet("vm-list", flag.ExitOnError)

		spinApiserverVMGetFlags  = flag.NewFlagSet("vm-get", flag.ExitOnError)
		spinApiserverVMGetIDFlag = spinApiserverVMGetFlags.String("id", "REQUIRED", "ID of VM to retrieve")

		spinApiserverVMUpdateFlags    = flag.NewFlagSet("vm-update", flag.ExitOnError)
		spinApiserverVMUpdateBodyFlag = spinApiserverVMUpdateFlags.String("body", "REQUIRED", "")
		spinApiserverVMUpdateIDFlag   = spinApiserverVMUpdateFlags.String("id", "REQUIRED", "ID of VM to Update")

		spinApiserverControlStartFlags  = flag.NewFlagSet("control-start", flag.ExitOnError)
		spinApiserverControlStartIDFlag = spinApiserverControlStartFlags.String("id", "REQUIRED", "ID of VM to start")

		spinApiserverControlStopFlags  = flag.NewFlagSet("control-stop", flag.ExitOnError)
		spinApiserverControlStopIDFlag = spinApiserverControlStopFlags.String("id", "REQUIRED", "ID of VM to stop")

		spinApiserverControlShutdownFlags  = flag.NewFlagSet("control-shutdown", flag.ExitOnError)
		spinApiserverControlShutdownIDFlag = spinApiserverControlShutdownFlags.String("id", "REQUIRED", "ID of VM to shutdown")

		spinRegistryFlags = flag.NewFlagSet("spin-registry", flag.ContinueOnError)

		spinRegistryVMCreateFlags    = flag.NewFlagSet("vm-create", flag.ExitOnError)
		spinRegistryVMCreateBodyFlag = spinRegistryVMCreateFlags.String("body", "REQUIRED", "")

		spinRegistryVMUpdateFlags    = flag.NewFlagSet("vm-update", flag.ExitOnError)
		spinRegistryVMUpdateBodyFlag = spinRegistryVMUpdateFlags.String("body", "REQUIRED", "")
		spinRegistryVMUpdateIDFlag   = spinRegistryVMUpdateFlags.String("id", "REQUIRED", "ID of VM to update")

		spinRegistryVMDeleteFlags  = flag.NewFlagSet("vm-delete", flag.ExitOnError)
		spinRegistryVMDeleteIDFlag = spinRegistryVMDeleteFlags.String("id", "REQUIRED", "ID of VM to remove")

		spinRegistryVMGetFlags  = flag.NewFlagSet("vm-get", flag.ExitOnError)
		spinRegistryVMGetIDFlag = spinRegistryVMGetFlags.String("id", "REQUIRED", "ID of VM to remove")

		spinRegistryVMListFlags = flag.NewFlagSet("vm-list", flag.ExitOnError)

		spinRegistryStorageVolumesListFlags = flag.NewFlagSet("storage-volumes-list", flag.ExitOnError)

		spinRegistryStorageVolumesCreateFlags    = flag.NewFlagSet("storage-volumes-create", flag.ExitOnError)
		spinRegistryStorageVolumesCreateBodyFlag = spinRegistryStorageVolumesCreateFlags.String("body", "REQUIRED", "")

		spinRegistryStorageVolumesDeleteFlags    = flag.NewFlagSet("storage-volumes-delete", flag.ExitOnError)
		spinRegistryStorageVolumesDeleteBodyFlag = spinRegistryStorageVolumesDeleteFlags.String("body", "REQUIRED", "")

		spinRegistryStorageImagesListFlags    = flag.NewFlagSet("storage-images-list", flag.ExitOnError)
		spinRegistryStorageImagesListBodyFlag = spinRegistryStorageImagesListFlags.String("body", "REQUIRED", "")

		spinRegistryStorageImagesCreateFlags    = flag.NewFlagSet("storage-images-create", flag.ExitOnError)
		spinRegistryStorageImagesCreateBodyFlag = spinRegistryStorageImagesCreateFlags.String("body", "REQUIRED", "")

		spinRegistryStorageImagesDeleteFlags    = flag.NewFlagSet("storage-images-delete", flag.ExitOnError)
		spinRegistryStorageImagesDeleteBodyFlag = spinRegistryStorageImagesDeleteFlags.String("body", "REQUIRED", "")

		spinRegistryStorageImagesGetFlags    = flag.NewFlagSet("storage-images-get", flag.ExitOnError)
		spinRegistryStorageImagesGetBodyFlag = spinRegistryStorageImagesGetFlags.String("body", "REQUIRED", "")
	)
	spinBrokerFlags.Usage = spinBrokerUsage
	spinBrokerNewFlags.Usage = spinBrokerNewUsage
	spinBrokerAddFlags.Usage = spinBrokerAddUsage
	spinBrokerEnqueueFlags.Usage = spinBrokerEnqueueUsage
	spinBrokerStatusFlags.Usage = spinBrokerStatusUsage
	spinBrokerNextFlags.Usage = spinBrokerNextUsage
	spinBrokerCompleteFlags.Usage = spinBrokerCompleteUsage

	spinApiserverFlags.Usage = spinApiserverUsage
	spinApiserverVMCreateFlags.Usage = spinApiserverVMCreateUsage
	spinApiserverVMDeleteFlags.Usage = spinApiserverVMDeleteUsage
	spinApiserverVMListFlags.Usage = spinApiserverVMListUsage
	spinApiserverVMGetFlags.Usage = spinApiserverVMGetUsage
	spinApiserverVMUpdateFlags.Usage = spinApiserverVMUpdateUsage
	spinApiserverControlStartFlags.Usage = spinApiserverControlStartUsage
	spinApiserverControlStopFlags.Usage = spinApiserverControlStopUsage
	spinApiserverControlShutdownFlags.Usage = spinApiserverControlShutdownUsage

	spinRegistryFlags.Usage = spinRegistryUsage
	spinRegistryVMCreateFlags.Usage = spinRegistryVMCreateUsage
	spinRegistryVMUpdateFlags.Usage = spinRegistryVMUpdateUsage
	spinRegistryVMDeleteFlags.Usage = spinRegistryVMDeleteUsage
	spinRegistryVMGetFlags.Usage = spinRegistryVMGetUsage
	spinRegistryVMListFlags.Usage = spinRegistryVMListUsage
	spinRegistryStorageVolumesListFlags.Usage = spinRegistryStorageVolumesListUsage
	spinRegistryStorageVolumesCreateFlags.Usage = spinRegistryStorageVolumesCreateUsage
	spinRegistryStorageVolumesDeleteFlags.Usage = spinRegistryStorageVolumesDeleteUsage
	spinRegistryStorageImagesListFlags.Usage = spinRegistryStorageImagesListUsage
	spinRegistryStorageImagesCreateFlags.Usage = spinRegistryStorageImagesCreateUsage
	spinRegistryStorageImagesDeleteFlags.Usage = spinRegistryStorageImagesDeleteUsage
	spinRegistryStorageImagesGetFlags.Usage = spinRegistryStorageImagesGetUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "spin-broker":
			svcf = spinBrokerFlags
		case "spin-apiserver":
			svcf = spinApiserverFlags
		case "spin-registry":
			svcf = spinRegistryFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "spin-broker":
			switch epn {
			case "new":
				epf = spinBrokerNewFlags

			case "add":
				epf = spinBrokerAddFlags

			case "enqueue":
				epf = spinBrokerEnqueueFlags

			case "status":
				epf = spinBrokerStatusFlags

			case "next":
				epf = spinBrokerNextFlags

			case "complete":
				epf = spinBrokerCompleteFlags

			}

		case "spin-apiserver":
			switch epn {
			case "vm-create":
				epf = spinApiserverVMCreateFlags

			case "vm-delete":
				epf = spinApiserverVMDeleteFlags

			case "vm-list":
				epf = spinApiserverVMListFlags

			case "vm-get":
				epf = spinApiserverVMGetFlags

			case "vm-update":
				epf = spinApiserverVMUpdateFlags

			case "control-start":
				epf = spinApiserverControlStartFlags

			case "control-stop":
				epf = spinApiserverControlStopFlags

			case "control-shutdown":
				epf = spinApiserverControlShutdownFlags

			}

		case "spin-registry":
			switch epn {
			case "vm-create":
				epf = spinRegistryVMCreateFlags

			case "vm-update":
				epf = spinRegistryVMUpdateFlags

			case "vm-delete":
				epf = spinRegistryVMDeleteFlags

			case "vm-get":
				epf = spinRegistryVMGetFlags

			case "vm-list":
				epf = spinRegistryVMListFlags

			case "storage-volumes-list":
				epf = spinRegistryStorageVolumesListFlags

			case "storage-volumes-create":
				epf = spinRegistryStorageVolumesCreateFlags

			case "storage-volumes-delete":
				epf = spinRegistryStorageVolumesDeleteFlags

			case "storage-images-list":
				epf = spinRegistryStorageImagesListFlags

			case "storage-images-create":
				epf = spinRegistryStorageImagesCreateFlags

			case "storage-images-delete":
				epf = spinRegistryStorageImagesDeleteFlags

			case "storage-images-get":
				epf = spinRegistryStorageImagesGetFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "spin-broker":
			c := spinbrokerc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "new":
				endpoint = c.New()
				data = nil
			case "add":
				endpoint = c.Add()
				data, err = spinbrokerc.BuildAddPayload(*spinBrokerAddBodyFlag, *spinBrokerAddIDFlag)
			case "enqueue":
				endpoint = c.Enqueue()
				data, err = spinbrokerc.BuildEnqueuePayload(*spinBrokerEnqueueIDFlag)
			case "status":
				endpoint = c.Status()
				data, err = spinbrokerc.BuildStatusPayload(*spinBrokerStatusIDFlag)
			case "next":
				endpoint = c.Next()
				data, err = spinbrokerc.BuildNextPayload(*spinBrokerNextResourceFlag)
			case "complete":
				endpoint = c.Complete()
				data, err = spinbrokerc.BuildCompletePayload(*spinBrokerCompleteBodyFlag)
			}
		case "spin-apiserver":
			c := spinapiserverc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "vm-create":
				endpoint = c.VMCreate()
				data, err = spinapiserverc.BuildVMCreatePayload(*spinApiserverVMCreateBodyFlag)
			case "vm-delete":
				endpoint = c.VMDelete()
				data, err = spinapiserverc.BuildVMDeletePayload(*spinApiserverVMDeleteIDFlag)
			case "vm-list":
				endpoint = c.VMList()
				data = nil
			case "vm-get":
				endpoint = c.VMGet()
				data, err = spinapiserverc.BuildVMGetPayload(*spinApiserverVMGetIDFlag)
			case "vm-update":
				endpoint = c.VMUpdate()
				data, err = spinapiserverc.BuildVMUpdatePayload(*spinApiserverVMUpdateBodyFlag, *spinApiserverVMUpdateIDFlag)
			case "control-start":
				endpoint = c.ControlStart()
				data, err = spinapiserverc.BuildControlStartPayload(*spinApiserverControlStartIDFlag)
			case "control-stop":
				endpoint = c.ControlStop()
				data, err = spinapiserverc.BuildControlStopPayload(*spinApiserverControlStopIDFlag)
			case "control-shutdown":
				endpoint = c.ControlShutdown()
				data, err = spinapiserverc.BuildControlShutdownPayload(*spinApiserverControlShutdownIDFlag)
			}
		case "spin-registry":
			c := spinregistryc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "vm-create":
				endpoint = c.VMCreate()
				data, err = spinregistryc.BuildVMCreatePayload(*spinRegistryVMCreateBodyFlag)
			case "vm-update":
				endpoint = c.VMUpdate()
				data, err = spinregistryc.BuildVMUpdatePayload(*spinRegistryVMUpdateBodyFlag, *spinRegistryVMUpdateIDFlag)
			case "vm-delete":
				endpoint = c.VMDelete()
				data, err = spinregistryc.BuildVMDeletePayload(*spinRegistryVMDeleteIDFlag)
			case "vm-get":
				endpoint = c.VMGet()
				data, err = spinregistryc.BuildVMGetPayload(*spinRegistryVMGetIDFlag)
			case "vm-list":
				endpoint = c.VMList()
				data = nil
			case "storage-volumes-list":
				endpoint = c.StorageVolumesList()
				data = nil
			case "storage-volumes-create":
				endpoint = c.StorageVolumesCreate()
				data, err = spinregistryc.BuildStorageVolumesCreatePayload(*spinRegistryStorageVolumesCreateBodyFlag)
			case "storage-volumes-delete":
				endpoint = c.StorageVolumesDelete()
				data, err = spinregistryc.BuildStorageVolumesDeletePayload(*spinRegistryStorageVolumesDeleteBodyFlag)
			case "storage-images-list":
				endpoint = c.StorageImagesList()
				data, err = spinregistryc.BuildStorageImagesListPayload(*spinRegistryStorageImagesListBodyFlag)
			case "storage-images-create":
				endpoint = c.StorageImagesCreate()
				data, err = spinregistryc.BuildStorageImagesCreatePayload(*spinRegistryStorageImagesCreateBodyFlag)
			case "storage-images-delete":
				endpoint = c.StorageImagesDelete()
				data, err = spinregistryc.BuildStorageImagesDeletePayload(*spinRegistryStorageImagesDeleteBodyFlag)
			case "storage-images-get":
				endpoint = c.StorageImagesGet()
				data, err = spinregistryc.BuildStorageImagesGetPayload(*spinRegistryStorageImagesGetBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// spin-brokerUsage displays the usage of the spin-broker command and its
// subcommands.
func spinBrokerUsage() {
	fmt.Fprintf(os.Stderr, `The message broker for the other services
Usage:
    %s [globalflags] spin-broker COMMAND [flags]

COMMAND:
    new: Create a new package; a collection of items to join into the queue simultaneously
    add: Add a command to the package
    enqueue: Enqueue the package into the various resource queues
    status: Get the status for a package
    next: Get the next command for a given resource
    complete: Mark a command as completed with a result status

Additional help:
    %s spin-broker COMMAND --help
`, os.Args[0], os.Args[0])
}
func spinBrokerNewUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker new

Create a new package; a collection of items to join into the queue simultaneously

Example:
    `+os.Args[0]+` spin-broker new
`, os.Args[0])
}

func spinBrokerAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker add -body JSON -id STRING

Add a command to the package
    -body JSON: 
    -id STRING: Package ID

Example:
    `+os.Args[0]+` spin-broker add --body '{
      "action": "Saepe occaecati impedit quas laborum.",
      "dependencies": [
         "Fugiat ex consequatur.",
         "Dignissimos eum.",
         "Velit quod sit aut."
      ],
      "parameters": {
         "Itaque aut dolorem.": "Quibusdam dolor sit."
      },
      "resource": "Id et autem ut debitis."
   }' --id "Natus temporibus fugit occaecati ipsum qui."
`, os.Args[0])
}

func spinBrokerEnqueueUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker enqueue -id STRING

Enqueue the package into the various resource queues
    -id STRING: Package ID

Example:
    `+os.Args[0]+` spin-broker enqueue --id "Voluptatibus nostrum commodi error omnis quis quia."
`, os.Args[0])
}

func spinBrokerStatusUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker status -id STRING

Get the status for a package
    -id STRING: Package ID

Example:
    `+os.Args[0]+` spin-broker status --id "Quidem qui."
`, os.Args[0])
}

func spinBrokerNextUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker next -resource STRING

Get the next command for a given resource
    -resource STRING: resource type

Example:
    `+os.Args[0]+` spin-broker next --resource "Eos iste illum omnis suscipit."
`, os.Args[0])
}

func spinBrokerCompleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker complete -body JSON

Mark a command as completed with a result status
    -body JSON: 

Example:
    `+os.Args[0]+` spin-broker complete --body '{
      "id": "Non sit recusandae eum repellat earum.",
      "status": false,
      "status_reason": "Eveniet hic minus facere animi quas tempora."
   }'
`, os.Args[0])
}

// spin-apiserverUsage displays the usage of the spin-apiserver command and its
// subcommands.
func spinApiserverUsage() {
	fmt.Fprintf(os.Stderr, `Bridge between the outer-facing UIs and the internals
Usage:
    %s [globalflags] spin-apiserver COMMAND [flags]

COMMAND:
    vm-create: VMCreate implements vm_create.
    vm-delete: VMDelete implements vm_delete.
    vm-list: VMList implements vm_list.
    vm-get: VMGet implements vm_get.
    vm-update: VMUpdate implements vm_update.
    control-start: ControlStart implements control_start.
    control-stop: ControlStop implements control_stop.
    control-shutdown: ControlShutdown implements control_shutdown.

Additional help:
    %s spin-apiserver COMMAND --help
`, os.Args[0], os.Args[0])
}
func spinApiserverVMCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver vm-create -body JSON

VMCreate implements vm_create.
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver vm-create --body '{
      "cpus": 17092856392700368829,
      "memory": 497222928895641055,
      "name": "Aspernatur vel vel illum voluptatem voluptatibus est.",
      "storage": [
         {
            "cdrom": true,
            "image": "Esse reprehenderit qui molestias eum voluptatem.",
            "image_size": 11533824901793082147,
            "volume": "Occaecati deserunt qui praesentium."
         },
         {
            "cdrom": true,
            "image": "Esse reprehenderit qui molestias eum voluptatem.",
            "image_size": 11533824901793082147,
            "volume": "Occaecati deserunt qui praesentium."
         }
      ]
   }'
`, os.Args[0])
}

func spinApiserverVMDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver vm-delete -id UINT64

VMDelete implements vm_delete.
    -id UINT64: ID of VM to delete

Example:
    `+os.Args[0]+` spin-apiserver vm-delete --id 17947092167867694340
`, os.Args[0])
}

func spinApiserverVMListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver vm-list

VMList implements vm_list.

Example:
    `+os.Args[0]+` spin-apiserver vm-list
`, os.Args[0])
}

func spinApiserverVMGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver vm-get -id UINT64

VMGet implements vm_get.
    -id UINT64: ID of VM to retrieve

Example:
    `+os.Args[0]+` spin-apiserver vm-get --id 2453152901843734876
`, os.Args[0])
}

func spinApiserverVMUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver vm-update -body JSON -id UINT64

VMUpdate implements vm_update.
    -body JSON: 
    -id UINT64: ID of VM to Update

Example:
    `+os.Args[0]+` spin-apiserver vm-update --body '{
      "vm": {
         "cpus": 3472743333644681302,
         "images": [
            {
               "cdrom": true,
               "path": "Dignissimos qui error modi.",
               "volume": "Corrupti et voluptatibus et et occaecati."
            },
            {
               "cdrom": true,
               "path": "Dignissimos qui error modi.",
               "volume": "Corrupti et voluptatibus et et occaecati."
            },
            {
               "cdrom": true,
               "path": "Dignissimos qui error modi.",
               "volume": "Corrupti et voluptatibus et et occaecati."
            }
         ],
         "memory": 639108202290023137,
         "name": "Quo dolore soluta consectetur."
      }
   }' --id 16411754489742442442
`, os.Args[0])
}

func spinApiserverControlStartUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver control-start -id UINT64

ControlStart implements control_start.
    -id UINT64: ID of VM to start

Example:
    `+os.Args[0]+` spin-apiserver control-start --id 3022875047415398183
`, os.Args[0])
}

func spinApiserverControlStopUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver control-stop -id UINT64

ControlStop implements control_stop.
    -id UINT64: ID of VM to stop

Example:
    `+os.Args[0]+` spin-apiserver control-stop --id 11176502060165847513
`, os.Args[0])
}

func spinApiserverControlShutdownUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver control-shutdown -id UINT64

ControlShutdown implements control_shutdown.
    -id UINT64: ID of VM to shutdown

Example:
    `+os.Args[0]+` spin-apiserver control-shutdown --id 12308344265677022532
`, os.Args[0])
}

// spin-registryUsage displays the usage of the spin-registry command and its
// subcommands.
func spinRegistryUsage() {
	fmt.Fprintf(os.Stderr, `Keeper of the VMs
Usage:
    %s [globalflags] spin-registry COMMAND [flags]

COMMAND:
    vm-create: Create a VM
    vm-update: Update a VM
    vm-delete: Delete a VM by ID
    vm-get: Retrieve a VM by ID
    vm-list: Retrieve all VM IDs
    storage-volumes-list: list all volumes
    storage-volumes-create: create a new volume
    storage-volumes-delete: delete an existing volume
    storage-images-list: list all images by volume
    storage-images-create: add an image definition to the registry
    storage-images-delete: remove an image definition from the registry
    storage-images-get: retrieves an image definition from the registry

Additional help:
    %s spin-registry COMMAND --help
`, os.Args[0], os.Args[0])
}
func spinRegistryVMCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-create -body JSON

Create a VM
    -body JSON: 

Example:
    `+os.Args[0]+` spin-registry vm-create --body '{
      "cpus": 2551606264180670795,
      "images": [
         {
            "cdrom": true,
            "path": "Dignissimos qui error modi.",
            "volume": "Corrupti et voluptatibus et et occaecati."
         },
         {
            "cdrom": true,
            "path": "Dignissimos qui error modi.",
            "volume": "Corrupti et voluptatibus et et occaecati."
         },
         {
            "cdrom": true,
            "path": "Dignissimos qui error modi.",
            "volume": "Corrupti et voluptatibus et et occaecati."
         },
         {
            "cdrom": true,
            "path": "Dignissimos qui error modi.",
            "volume": "Corrupti et voluptatibus et et occaecati."
         }
      ],
      "memory": 17694085759044319548,
      "name": "Nobis quia."
   }'
`, os.Args[0])
}

func spinRegistryVMUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-update -body JSON -id UINT64

Update a VM
    -body JSON: 
    -id UINT64: ID of VM to update

Example:
    `+os.Args[0]+` spin-registry vm-update --body '{
      "vm": {
         "cpus": 3472743333644681302,
         "images": [
            {
               "cdrom": true,
               "path": "Dignissimos qui error modi.",
               "volume": "Corrupti et voluptatibus et et occaecati."
            },
            {
               "cdrom": true,
               "path": "Dignissimos qui error modi.",
               "volume": "Corrupti et voluptatibus et et occaecati."
            },
            {
               "cdrom": true,
               "path": "Dignissimos qui error modi.",
               "volume": "Corrupti et voluptatibus et et occaecati."
            }
         ],
         "memory": 639108202290023137,
         "name": "Quo dolore soluta consectetur."
      }
   }' --id 13160086984666010830
`, os.Args[0])
}

func spinRegistryVMDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-delete -id UINT64

Delete a VM by ID
    -id UINT64: ID of VM to remove

Example:
    `+os.Args[0]+` spin-registry vm-delete --id 12531413382890778925
`, os.Args[0])
}

func spinRegistryVMGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-get -id UINT64

Retrieve a VM by ID
    -id UINT64: ID of VM to remove

Example:
    `+os.Args[0]+` spin-registry vm-get --id 1456500719710708694
`, os.Args[0])
}

func spinRegistryVMListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-list

Retrieve all VM IDs

Example:
    `+os.Args[0]+` spin-registry vm-list
`, os.Args[0])
}

func spinRegistryStorageVolumesListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry storage-volumes-list

list all volumes

Example:
    `+os.Args[0]+` spin-registry storage-volumes-list
`, os.Args[0])
}

func spinRegistryStorageVolumesCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry storage-volumes-create -body JSON

create a new volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-registry storage-volumes-create --body '{
      "name": "Ut ea excepturi.",
      "path": "Facilis ad quod."
   }'
`, os.Args[0])
}

func spinRegistryStorageVolumesDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry storage-volumes-delete -body JSON

delete an existing volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-registry storage-volumes-delete --body '{
      "name": "Quia facere."
   }'
`, os.Args[0])
}

func spinRegistryStorageImagesListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry storage-images-list -body JSON

list all images by volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-registry storage-images-list --body '{
      "volume_name": "Sunt nesciunt natus dolorem."
   }'
`, os.Args[0])
}

func spinRegistryStorageImagesCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry storage-images-create -body JSON

add an image definition to the registry
    -body JSON: 

Example:
    `+os.Args[0]+` spin-registry storage-images-create --body '{
      "cdrom": true,
      "image": "Nostrum qui perferendis rerum molestias.",
      "image_size": 11387401285549308839,
      "volume": "Iure qui voluptas."
   }'
`, os.Args[0])
}

func spinRegistryStorageImagesDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry storage-images-delete -body JSON

remove an image definition from the registry
    -body JSON: 

Example:
    `+os.Args[0]+` spin-registry storage-images-delete --body '{
      "image_name": "Ipsam dicta accusantium.",
      "volume_name": "Impedit laboriosam et dolorum tempora inventore officia."
   }'
`, os.Args[0])
}

func spinRegistryStorageImagesGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry storage-images-get -body JSON

retrieves an image definition from the registry
    -body JSON: 

Example:
    `+os.Args[0]+` spin-registry storage-images-get --body '{
      "image_name": "Laudantium consectetur assumenda soluta.",
      "volume_name": "In ea natus tempore."
   }'
`, os.Args[0])
}
