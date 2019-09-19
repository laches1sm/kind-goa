// Code generated by goa v3.0.6, DO NOT EDIT.
//
// parrot HTTP client CLI support package
//
// Command:
// $ goa gen parrot_service/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	healthc "parrot_service/gen/http/health/client"
	serviceparrotc "parrot_service/gen/http/service_parrot/client"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `health check
service-parrot (postparrot|listaparrot|listallparrots)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` health check` + "\n" +
		os.Args[0] + ` service-parrot postparrot` + "\n" +
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
		healthFlags = flag.NewFlagSet("health", flag.ContinueOnError)

		healthCheckFlags = flag.NewFlagSet("check", flag.ExitOnError)

		serviceParrotFlags = flag.NewFlagSet("service-parrot", flag.ContinueOnError)

		serviceParrotPostparrotFlags = flag.NewFlagSet("postparrot", flag.ExitOnError)

		serviceParrotListaparrotFlags    = flag.NewFlagSet("listaparrot", flag.ExitOnError)
		serviceParrotListaparrotBodyFlag = serviceParrotListaparrotFlags.String("body", "REQUIRED", "")

		serviceParrotListallparrotsFlags = flag.NewFlagSet("listallparrots", flag.ExitOnError)
	)
	healthFlags.Usage = healthUsage
	healthCheckFlags.Usage = healthCheckUsage

	serviceParrotFlags.Usage = serviceParrotUsage
	serviceParrotPostparrotFlags.Usage = serviceParrotPostparrotUsage
	serviceParrotListaparrotFlags.Usage = serviceParrotListaparrotUsage
	serviceParrotListallparrotsFlags.Usage = serviceParrotListallparrotsUsage

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
		case "health":
			svcf = healthFlags
		case "service-parrot":
			svcf = serviceParrotFlags
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
		case "health":
			switch epn {
			case "check":
				epf = healthCheckFlags

			}

		case "service-parrot":
			switch epn {
			case "postparrot":
				epf = serviceParrotPostparrotFlags

			case "listaparrot":
				epf = serviceParrotListaparrotFlags

			case "listallparrots":
				epf = serviceParrotListallparrotsFlags

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
		case "health":
			c := healthc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "check":
				endpoint = c.Check()
				data = nil
			}
		case "service-parrot":
			c := serviceparrotc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "postparrot":
				endpoint = c.Postparrot()
				data, err = serviceparrotc.BuildPostparrotPayload()
			case "listaparrot":
				endpoint = c.Listaparrot()
				data, err = serviceparrotc.BuildListaparrotPayload(*serviceParrotListaparrotBodyFlag)
			case "listallparrots":
				endpoint = c.Listallparrots()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// healthUsage displays the usage of the health command and its subcommands.
func healthUsage() {
	fmt.Fprintf(os.Stderr, `Service is the Health service interface.
Usage:
    %s [globalflags] health COMMAND [flags]

COMMAND:
    check: Health check

Additional help:
    %s health COMMAND --help
`, os.Args[0], os.Args[0])
}
func healthCheckUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] health check

Health check

Example:
    `+os.Args[0]+` health check
`, os.Args[0])
}

// service-parrotUsage displays the usage of the service-parrot command and its
// subcommands.
func serviceParrotUsage() {
	fmt.Fprintf(os.Stderr, `The Parrot service serves a parrot payload
Usage:
    %s [globalflags] service-parrot COMMAND [flags]

COMMAND:
    postparrot: Postparrot implements postparrot.
    listaparrot: Listaparrot implements listaparrot.
    listallparrots: Listallparrots implements listallparrots.

Additional help:
    %s service-parrot COMMAND --help
`, os.Args[0], os.Args[0])
}
func serviceParrotPostparrotUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] service-parrot postparrot

Postparrot implements postparrot.

Example:
    `+os.Args[0]+` service-parrot postparrot
`, os.Args[0])
}

func serviceParrotListaparrotUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] service-parrot listaparrot -body JSON

Listaparrot implements listaparrot.
    -body JSON: 

Example:
    `+os.Args[0]+` service-parrot listaparrot --body '{
      "id": "123"
   }'
`, os.Args[0])
}

func serviceParrotListallparrotsUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] service-parrot listallparrots

Listallparrots implements listallparrots.

Example:
    `+os.Args[0]+` service-parrot listallparrots
`, os.Args[0])
}
