// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin HTTP client CLI support package
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	spinapiserverc "code.hollensbe.org/erikh/spin/gen/http/spin_apiserver/client"
	spinbrokerc "code.hollensbe.org/erikh/spin/gen/http/spin_broker/client"
	spinregistryc "code.hollensbe.org/erikh/spin/gen/http/spin_registry/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `spin-apiserver (add-volume|remove-volume|info-volume|create-image-on-volume|delete-image-on-volume|resize-image-on-volume|info-image-on-volume|move-image)
spin-broker (new|add|enqueue|status|next|complete)
spin-registry (vm-/create|vm-/update|vm-/delete|vm-/get|vm-/list)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` spin-apiserver add-volume --body '{
      "path": "Eum officiis eligendi.",
      "volume": "Non deleniti consequuntur qui doloremque."
   }'` + "\n" +
		os.Args[0] + ` spin-broker new` + "\n" +
		os.Args[0] + ` spin-registry vm-/create --body '{
      "cpus": 6421225864845588134,
      "memory": 15540618699971399982,
      "name": "Et dicta.",
      "storage": [
         {
            "cdrom": true,
            "image": "Labore voluptas perferendis ea iusto adipisci.",
            "image_size": 8351098711286704476,
            "volume": "Porro eius officiis."
         },
         {
            "cdrom": true,
            "image": "Labore voluptas perferendis ea iusto adipisci.",
            "image_size": 8351098711286704476,
            "volume": "Porro eius officiis."
         },
         {
            "cdrom": true,
            "image": "Labore voluptas perferendis ea iusto adipisci.",
            "image_size": 8351098711286704476,
            "volume": "Porro eius officiis."
         },
         {
            "cdrom": true,
            "image": "Labore voluptas perferendis ea iusto adipisci.",
            "image_size": 8351098711286704476,
            "volume": "Porro eius officiis."
         }
      ]
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
		spinApiserverFlags = flag.NewFlagSet("spin-apiserver", flag.ContinueOnError)

		spinApiserverAddVolumeFlags    = flag.NewFlagSet("add-volume", flag.ExitOnError)
		spinApiserverAddVolumeBodyFlag = spinApiserverAddVolumeFlags.String("body", "REQUIRED", "")

		spinApiserverRemoveVolumeFlags      = flag.NewFlagSet("remove-volume", flag.ExitOnError)
		spinApiserverRemoveVolumeVolumeFlag = spinApiserverRemoveVolumeFlags.String("volume", "REQUIRED", "volume identifier")

		spinApiserverInfoVolumeFlags      = flag.NewFlagSet("info-volume", flag.ExitOnError)
		spinApiserverInfoVolumeVolumeFlag = spinApiserverInfoVolumeFlags.String("volume", "REQUIRED", "volume identifier")

		spinApiserverCreateImageOnVolumeFlags    = flag.NewFlagSet("create-image-on-volume", flag.ExitOnError)
		spinApiserverCreateImageOnVolumeBodyFlag = spinApiserverCreateImageOnVolumeFlags.String("body", "REQUIRED", "")

		spinApiserverDeleteImageOnVolumeFlags    = flag.NewFlagSet("delete-image-on-volume", flag.ExitOnError)
		spinApiserverDeleteImageOnVolumeBodyFlag = spinApiserverDeleteImageOnVolumeFlags.String("body", "REQUIRED", "")

		spinApiserverResizeImageOnVolumeFlags    = flag.NewFlagSet("resize-image-on-volume", flag.ExitOnError)
		spinApiserverResizeImageOnVolumeBodyFlag = spinApiserverResizeImageOnVolumeFlags.String("body", "REQUIRED", "")

		spinApiserverInfoImageOnVolumeFlags         = flag.NewFlagSet("info-image-on-volume", flag.ExitOnError)
		spinApiserverInfoImageOnVolumeVolumeFlag    = spinApiserverInfoImageOnVolumeFlags.String("volume", "REQUIRED", "volume identifier")
		spinApiserverInfoImageOnVolumeImageNameFlag = spinApiserverInfoImageOnVolumeFlags.String("image-name", "REQUIRED", "image name")

		spinApiserverMoveImageFlags    = flag.NewFlagSet("move-image", flag.ExitOnError)
		spinApiserverMoveImageBodyFlag = spinApiserverMoveImageFlags.String("body", "REQUIRED", "")

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

		spinRegistryFlags = flag.NewFlagSet("spin-registry", flag.ContinueOnError)

		spinRegistryVMCreateFlags    = flag.NewFlagSet("vm-/create", flag.ExitOnError)
		spinRegistryVMCreateBodyFlag = spinRegistryVMCreateFlags.String("body", "REQUIRED", "")

		spinRegistryVMUpdateFlags    = flag.NewFlagSet("vm-/update", flag.ExitOnError)
		spinRegistryVMUpdateBodyFlag = spinRegistryVMUpdateFlags.String("body", "REQUIRED", "")
		spinRegistryVMUpdateIDFlag   = spinRegistryVMUpdateFlags.String("id", "REQUIRED", "ID of VM to update")

		spinRegistryVMDeleteFlags  = flag.NewFlagSet("vm-/delete", flag.ExitOnError)
		spinRegistryVMDeleteIDFlag = spinRegistryVMDeleteFlags.String("id", "REQUIRED", "ID of VM to remove")

		spinRegistryVMGetFlags  = flag.NewFlagSet("vm-/get", flag.ExitOnError)
		spinRegistryVMGetIDFlag = spinRegistryVMGetFlags.String("id", "REQUIRED", "ID of VM to remove")

		spinRegistryVMListFlags = flag.NewFlagSet("vm-/list", flag.ExitOnError)
	)
	spinApiserverFlags.Usage = spinApiserverUsage
	spinApiserverAddVolumeFlags.Usage = spinApiserverAddVolumeUsage
	spinApiserverRemoveVolumeFlags.Usage = spinApiserverRemoveVolumeUsage
	spinApiserverInfoVolumeFlags.Usage = spinApiserverInfoVolumeUsage
	spinApiserverCreateImageOnVolumeFlags.Usage = spinApiserverCreateImageOnVolumeUsage
	spinApiserverDeleteImageOnVolumeFlags.Usage = spinApiserverDeleteImageOnVolumeUsage
	spinApiserverResizeImageOnVolumeFlags.Usage = spinApiserverResizeImageOnVolumeUsage
	spinApiserverInfoImageOnVolumeFlags.Usage = spinApiserverInfoImageOnVolumeUsage
	spinApiserverMoveImageFlags.Usage = spinApiserverMoveImageUsage

	spinBrokerFlags.Usage = spinBrokerUsage
	spinBrokerNewFlags.Usage = spinBrokerNewUsage
	spinBrokerAddFlags.Usage = spinBrokerAddUsage
	spinBrokerEnqueueFlags.Usage = spinBrokerEnqueueUsage
	spinBrokerStatusFlags.Usage = spinBrokerStatusUsage
	spinBrokerNextFlags.Usage = spinBrokerNextUsage
	spinBrokerCompleteFlags.Usage = spinBrokerCompleteUsage

	spinRegistryFlags.Usage = spinRegistryUsage
	spinRegistryVMCreateFlags.Usage = spinRegistryVMCreateUsage
	spinRegistryVMUpdateFlags.Usage = spinRegistryVMUpdateUsage
	spinRegistryVMDeleteFlags.Usage = spinRegistryVMDeleteUsage
	spinRegistryVMGetFlags.Usage = spinRegistryVMGetUsage
	spinRegistryVMListFlags.Usage = spinRegistryVMListUsage

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
		case "spin-apiserver":
			svcf = spinApiserverFlags
		case "spin-broker":
			svcf = spinBrokerFlags
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
		case "spin-apiserver":
			switch epn {
			case "add-volume":
				epf = spinApiserverAddVolumeFlags

			case "remove-volume":
				epf = spinApiserverRemoveVolumeFlags

			case "info-volume":
				epf = spinApiserverInfoVolumeFlags

			case "create-image-on-volume":
				epf = spinApiserverCreateImageOnVolumeFlags

			case "delete-image-on-volume":
				epf = spinApiserverDeleteImageOnVolumeFlags

			case "resize-image-on-volume":
				epf = spinApiserverResizeImageOnVolumeFlags

			case "info-image-on-volume":
				epf = spinApiserverInfoImageOnVolumeFlags

			case "move-image":
				epf = spinApiserverMoveImageFlags

			}

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

		case "spin-registry":
			switch epn {
			case "vm-/create":
				epf = spinRegistryVMCreateFlags

			case "vm-/update":
				epf = spinRegistryVMUpdateFlags

			case "vm-/delete":
				epf = spinRegistryVMDeleteFlags

			case "vm-/get":
				epf = spinRegistryVMGetFlags

			case "vm-/list":
				epf = spinRegistryVMListFlags

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
		case "spin-apiserver":
			c := spinapiserverc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add-volume":
				endpoint = c.AddVolume()
				data, err = spinapiserverc.BuildAddVolumePayload(*spinApiserverAddVolumeBodyFlag)
			case "remove-volume":
				endpoint = c.RemoveVolume()
				data, err = spinapiserverc.BuildRemoveVolumePayload(*spinApiserverRemoveVolumeVolumeFlag)
			case "info-volume":
				endpoint = c.InfoVolume()
				data, err = spinapiserverc.BuildInfoVolumePayload(*spinApiserverInfoVolumeVolumeFlag)
			case "create-image-on-volume":
				endpoint = c.CreateImageOnVolume()
				data, err = spinapiserverc.BuildCreateImageOnVolumePayload(*spinApiserverCreateImageOnVolumeBodyFlag)
			case "delete-image-on-volume":
				endpoint = c.DeleteImageOnVolume()
				data, err = spinapiserverc.BuildDeleteImageOnVolumePayload(*spinApiserverDeleteImageOnVolumeBodyFlag)
			case "resize-image-on-volume":
				endpoint = c.ResizeImageOnVolume()
				data, err = spinapiserverc.BuildResizeImageOnVolumePayload(*spinApiserverResizeImageOnVolumeBodyFlag)
			case "info-image-on-volume":
				endpoint = c.InfoImageOnVolume()
				data, err = spinapiserverc.BuildInfoImageOnVolumePayload(*spinApiserverInfoImageOnVolumeVolumeFlag, *spinApiserverInfoImageOnVolumeImageNameFlag)
			case "move-image":
				endpoint = c.MoveImage()
				data, err = spinapiserverc.BuildMoveImagePayload(*spinApiserverMoveImageBodyFlag)
			}
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
		case "spin-registry":
			c := spinregistryc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "vm-/create":
				endpoint = c.VMCreate()
				data, err = spinregistryc.BuildVMCreatePayload(*spinRegistryVMCreateBodyFlag)
			case "vm-/update":
				endpoint = c.VMUpdate()
				data, err = spinregistryc.BuildVMUpdatePayload(*spinRegistryVMUpdateBodyFlag, *spinRegistryVMUpdateIDFlag)
			case "vm-/delete":
				endpoint = c.VMDelete()
				data, err = spinregistryc.BuildVMDeletePayload(*spinRegistryVMDeleteIDFlag)
			case "vm-/get":
				endpoint = c.VMGet()
				data, err = spinregistryc.BuildVMGetPayload(*spinRegistryVMGetIDFlag)
			case "vm-/list":
				endpoint = c.VMList()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// spin-apiserverUsage displays the usage of the spin-apiserver command and its
// subcommands.
func spinApiserverUsage() {
	fmt.Fprintf(os.Stderr, `Bridge between the outer-facing UIs and the internals
Usage:
    %s [globalflags] spin-apiserver COMMAND [flags]

COMMAND:
    add-volume: Add a volume for image allocation with backing storage, and name it
    remove-volume: Remove a volume. Requires all images to be removed.
    info-volume: Get information on a volume
    create-image-on-volume: Create an image on a volume
    delete-image-on-volume: Delete an image on a volume
    resize-image-on-volume: Resize an image on a volume
    info-image-on-volume: Obtain information on an image
    move-image: Move an image from one volume to another

Additional help:
    %s spin-apiserver COMMAND --help
`, os.Args[0], os.Args[0])
}
func spinApiserverAddVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver add-volume -body JSON

Add a volume for image allocation with backing storage, and name it
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver add-volume --body '{
      "path": "Eum officiis eligendi.",
      "volume": "Non deleniti consequuntur qui doloremque."
   }'
`, os.Args[0])
}

func spinApiserverRemoveVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver remove-volume -volume STRING

Remove a volume. Requires all images to be removed.
    -volume STRING: volume identifier

Example:
    `+os.Args[0]+` spin-apiserver remove-volume --volume "Distinctio magni quibusdam reiciendis ipsum commodi sed."
`, os.Args[0])
}

func spinApiserverInfoVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver info-volume -volume STRING

Get information on a volume
    -volume STRING: volume identifier

Example:
    `+os.Args[0]+` spin-apiserver info-volume --volume "Ipsam dolor distinctio."
`, os.Args[0])
}

func spinApiserverCreateImageOnVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver create-image-on-volume -body JSON

Create an image on a volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver create-image-on-volume --body '{
      "image_name": "Vitae sunt amet nostrum amet.",
      "image_size": 9785040533924627437,
      "volume": "Dolorem odio laborum et."
   }'
`, os.Args[0])
}

func spinApiserverDeleteImageOnVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver delete-image-on-volume -body JSON

Delete an image on a volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver delete-image-on-volume --body '{
      "image_name": "Rerum ut voluptatem at fugit fuga impedit.",
      "volume": "Veritatis quo non quae rerum officia iusto."
   }'
`, os.Args[0])
}

func spinApiserverResizeImageOnVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver resize-image-on-volume -body JSON

Resize an image on a volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver resize-image-on-volume --body '{
      "image_name": "Et labore.",
      "image_size": 8231368540121560477,
      "volume": "Consequatur omnis dolor assumenda."
   }'
`, os.Args[0])
}

func spinApiserverInfoImageOnVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver info-image-on-volume -volume STRING -image-name STRING

Obtain information on an image
    -volume STRING: volume identifier
    -image-name STRING: image name

Example:
    `+os.Args[0]+` spin-apiserver info-image-on-volume --volume "Quia voluptas aut." --image-name "Dicta quisquam."
`, os.Args[0])
}

func spinApiserverMoveImageUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver move-image -body JSON

Move an image from one volume to another
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver move-image --body '{
      "image_name": "Qui ut.",
      "target_volume": "Nesciunt omnis voluptatem ullam.",
      "volume": "Est reprehenderit."
   }'
`, os.Args[0])
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
      "action": "Id et autem ut debitis.",
      "parameters": {
         "Occaecati impedit quas laborum consequatur.": "Itaque aut dolorem.",
         "Possimus velit quod sit.": "Mollitia natus temporibus fugit occaecati ipsum qui.",
         "Quibusdam dolor sit.": "Porro fugiat ex consequatur provident dignissimos."
      },
      "resource": "Quam neque."
   }' --id "Non aut molestiae deleniti ad dolorem eos."
`, os.Args[0])
}

func spinBrokerEnqueueUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker enqueue -id STRING

Enqueue the package into the various resource queues
    -id STRING: Package ID

Example:
    `+os.Args[0]+` spin-broker enqueue --id "Ratione et ex tempore ipsum in."
`, os.Args[0])
}

func spinBrokerStatusUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker status -id STRING

Get the status for a package
    -id STRING: Package ID

Example:
    `+os.Args[0]+` spin-broker status --id "Iste illum omnis."
`, os.Args[0])
}

func spinBrokerNextUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker next -resource STRING

Get the next command for a given resource
    -resource STRING: resource type

Example:
    `+os.Args[0]+` spin-broker next --resource "Iure aut voluptas qui."
`, os.Args[0])
}

func spinBrokerCompleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker complete -body JSON

Mark a command as completed with a result status
    -body JSON: 

Example:
    `+os.Args[0]+` spin-broker complete --body '{
      "id": "Cum delectus occaecati enim in.",
      "status": true,
      "status_reason": "Amet quia quis omnis suscipit."
   }'
`, os.Args[0])
}

// spin-registryUsage displays the usage of the spin-registry command and its
// subcommands.
func spinRegistryUsage() {
	fmt.Fprintf(os.Stderr, `Keeper of the VMs
Usage:
    %s [globalflags] spin-registry COMMAND [flags]

COMMAND:
    vm-/create: Create a VM
    vm-/update: Update a VM
    vm-/delete: Delete a VM by ID
    vm-/get: Retrieve a VM by ID
    vm-/list: Retrieve all VM IDs

Additional help:
    %s spin-registry COMMAND --help
`, os.Args[0], os.Args[0])
}
func spinRegistryVMCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-/create -body JSON

Create a VM
    -body JSON: 

Example:
    `+os.Args[0]+` spin-registry vm-/create --body '{
      "cpus": 6421225864845588134,
      "memory": 15540618699971399982,
      "name": "Et dicta.",
      "storage": [
         {
            "cdrom": true,
            "image": "Labore voluptas perferendis ea iusto adipisci.",
            "image_size": 8351098711286704476,
            "volume": "Porro eius officiis."
         },
         {
            "cdrom": true,
            "image": "Labore voluptas perferendis ea iusto adipisci.",
            "image_size": 8351098711286704476,
            "volume": "Porro eius officiis."
         },
         {
            "cdrom": true,
            "image": "Labore voluptas perferendis ea iusto adipisci.",
            "image_size": 8351098711286704476,
            "volume": "Porro eius officiis."
         },
         {
            "cdrom": true,
            "image": "Labore voluptas perferendis ea iusto adipisci.",
            "image_size": 8351098711286704476,
            "volume": "Porro eius officiis."
         }
      ]
   }'
`, os.Args[0])
}

func spinRegistryVMUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-/update -body JSON -id UINT64

Update a VM
    -body JSON: 
    -id UINT64: ID of VM to update

Example:
    `+os.Args[0]+` spin-registry vm-/update --body '{
      "vm": {
         "cpus": 8415388083871949362,
         "memory": 15948619849308310515,
         "name": "Et qui fugit quis dignissimos qui.",
         "storage": [
            {
               "cdrom": true,
               "image": "Labore voluptas perferendis ea iusto adipisci.",
               "image_size": 8351098711286704476,
               "volume": "Porro eius officiis."
            },
            {
               "cdrom": true,
               "image": "Labore voluptas perferendis ea iusto adipisci.",
               "image_size": 8351098711286704476,
               "volume": "Porro eius officiis."
            },
            {
               "cdrom": true,
               "image": "Labore voluptas perferendis ea iusto adipisci.",
               "image_size": 8351098711286704476,
               "volume": "Porro eius officiis."
            }
         ]
      }
   }' --id 5309827991366441111
`, os.Args[0])
}

func spinRegistryVMDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-/delete -id UINT64

Delete a VM by ID
    -id UINT64: ID of VM to remove

Example:
    `+os.Args[0]+` spin-registry vm-/delete --id 6259074207733349285
`, os.Args[0])
}

func spinRegistryVMGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-/get -id UINT64

Retrieve a VM by ID
    -id UINT64: ID of VM to remove

Example:
    `+os.Args[0]+` spin-registry vm-/get --id 5430522408910716022
`, os.Args[0])
}

func spinRegistryVMListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-registry vm-/list

Retrieve all VM IDs

Example:
    `+os.Args[0]+` spin-registry vm-/list
`, os.Args[0])
}
