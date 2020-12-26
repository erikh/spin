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
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `spin-apiserver (add-volume|remove-volume|label-volume|info-volume|create-image-on-volume|delete-image-on-volume|resize-image-on-volume|info-image-on-volume|move-image)
spin-broker (new|add|enqueue|status|next|complete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` spin-apiserver add-volume --body '{
      "path": "Eum officiis eligendi.",
      "volume": "Non deleniti consequuntur qui doloremque."
   }'` + "\n" +
		os.Args[0] + ` spin-broker new` + "\n" +
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

		spinApiserverLabelVolumeFlags      = flag.NewFlagSet("label-volume", flag.ExitOnError)
		spinApiserverLabelVolumeVolumeFlag = spinApiserverLabelVolumeFlags.String("volume", "REQUIRED", "volume identifier")
		spinApiserverLabelVolumeLabelFlag  = spinApiserverLabelVolumeFlags.String("label", "REQUIRED", "label identifier to apply to volume")

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
	)
	spinApiserverFlags.Usage = spinApiserverUsage
	spinApiserverAddVolumeFlags.Usage = spinApiserverAddVolumeUsage
	spinApiserverRemoveVolumeFlags.Usage = spinApiserverRemoveVolumeUsage
	spinApiserverLabelVolumeFlags.Usage = spinApiserverLabelVolumeUsage
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

			case "label-volume":
				epf = spinApiserverLabelVolumeFlags

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
			case "label-volume":
				endpoint = c.LabelVolume()
				data, err = spinapiserverc.BuildLabelVolumePayload(*spinApiserverLabelVolumeVolumeFlag, *spinApiserverLabelVolumeLabelFlag)
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
    label-volume: Apply a label to a volume.
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

func spinApiserverLabelVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver label-volume -volume STRING -label STRING

Apply a label to a volume.
    -volume STRING: volume identifier
    -label STRING: label identifier to apply to volume

Example:
    `+os.Args[0]+` spin-apiserver label-volume --volume "Vitae sunt amet nostrum amet." --label "Itaque dolorem odio."
`, os.Args[0])
}

func spinApiserverInfoVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver info-volume -volume STRING

Get information on a volume
    -volume STRING: volume identifier

Example:
    `+os.Args[0]+` spin-apiserver info-volume --volume "Impedit dolor veritatis quo non quae rerum."
`, os.Args[0])
}

func spinApiserverCreateImageOnVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver create-image-on-volume -body JSON

Create an image on a volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver create-image-on-volume --body '{
      "image_name": "Iusto minima et labore.",
      "image_size": 8231368540121560477,
      "volume": "Consequatur omnis dolor assumenda."
   }'
`, os.Args[0])
}

func spinApiserverDeleteImageOnVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver delete-image-on-volume -body JSON

Delete an image on a volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver delete-image-on-volume --body '{
      "image_name": "Dignissimos ut doloremque.",
      "volume": "Et tempora rem."
   }'
`, os.Args[0])
}

func spinApiserverResizeImageOnVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver resize-image-on-volume -body JSON

Resize an image on a volume
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver resize-image-on-volume --body '{
      "image_name": "Quia voluptas aut.",
      "image_size": 4442355973768769894,
      "volume": "Quisquam rem."
   }'
`, os.Args[0])
}

func spinApiserverInfoImageOnVolumeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver info-image-on-volume -volume STRING -image-name STRING

Obtain information on an image
    -volume STRING: volume identifier
    -image-name STRING: image name

Example:
    `+os.Args[0]+` spin-apiserver info-image-on-volume --volume "Ullam amet magnam." --image-name "Ad ut nulla laboriosam non deserunt vitae."
`, os.Args[0])
}

func spinApiserverMoveImageUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-apiserver move-image -body JSON

Move an image from one volume to another
    -body JSON: 

Example:
    `+os.Args[0]+` spin-apiserver move-image --body '{
      "image_name": "Illum sint.",
      "target_volume": "Quam neque.",
      "volume": "Ullam odio tenetur aliquid consequatur."
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
    new: New implements new.
    add: Add implements add.
    enqueue: Enqueue implements enqueue.
    status: Status implements status.
    next: Next implements next.
    complete: Complete implements complete.

Additional help:
    %s spin-broker COMMAND --help
`, os.Args[0], os.Args[0])
}
func spinBrokerNewUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker new

New implements new.

Example:
    `+os.Args[0]+` spin-broker new
`, os.Args[0])
}

func spinBrokerAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker add -body JSON -id STRING

Add implements add.
    -body JSON: 
    -id STRING: Package ID

Example:
    `+os.Args[0]+` spin-broker add --body '{
      "action": "Officiis porro fugiat.",
      "parameters": [
         "Provident dignissimos eum possimus velit quod.",
         "Aut mollitia natus temporibus fugit occaecati.",
         "Qui dolorem non aut molestiae.",
         "Ad dolorem eos rerum quam adipisci est."
      ],
      "resource": "Voluptatem quibusdam dolor."
   }' --id "Sequi consequatur maxime aut non."
`, os.Args[0])
}

func spinBrokerEnqueueUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker enqueue -id STRING

Enqueue implements enqueue.
    -id STRING: Package ID

Example:
    `+os.Args[0]+` spin-broker enqueue --id "Consectetur et."
`, os.Args[0])
}

func spinBrokerStatusUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker status -id STRING

Status implements status.
    -id STRING: Package ID

Example:
    `+os.Args[0]+` spin-broker status --id "Natus non consequatur voluptatem nostrum."
`, os.Args[0])
}

func spinBrokerNextUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker next -resource STRING

Next implements next.
    -resource STRING: resource type

Example:
    `+os.Args[0]+` spin-broker next --resource "Neque reiciendis ipsum."
`, os.Args[0])
}

func spinBrokerCompleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] spin-broker complete -body JSON

Complete implements complete.
    -body JSON: 

Example:
    `+os.Args[0]+` spin-broker complete --body '{
      "id": "Quas tempora nisi aperiam occaecati deserunt qui.",
      "status": false,
      "status_reason": "Esse reprehenderit qui molestias eum voluptatem."
   }'
`, os.Args[0])
}
