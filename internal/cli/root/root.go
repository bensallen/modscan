package root

import (
	"errors"
	"fmt"
	"os"

	"github.com/bensallen/modscan/internal/cli/list"
	"github.com/bensallen/modscan/internal/cli/load"
	flag "github.com/spf13/pflag"
)

// Version is the CLI version or release number. To be overriden at build time.
var Version = "unknown"

const usageHeader = `modscan - kernel module detection based on modalias

Usage:
  modscan [load|list]

Subcommands:
  load         Load needed kernel modules
  [list|print] List needed kernel modules

Flags:
`

var (
	rootFlags = flag.NewFlagSet("root", flag.ContinueOnError)
	help      = rootFlags.BoolP("help", "h", false, "Diplay help.")
	version   = rootFlags.BoolP("version", "V", false, "Displays the program version string.")
	verbose   = rootFlags.BoolP("verbose", "v", false, "Enable additional output.")
)

//Usage of modscan command printed to stderr
func usage() {
	fmt.Fprintf(os.Stderr, usageHeader)
	fmt.Fprintf(os.Stderr, rootFlags.FlagUsagesWrapped(0)+"\n")
}

//Usage of modscan with error message
func usageErr(err error) {
	usage()
	fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
}

// Run modscan
func Run(args []string) error {

	rootFlags.ParseErrorsWhitelist.UnknownFlags = true
	if err := rootFlags.Parse(args); err != nil {
		usageErr(err)
		os.Exit(2)
	}

	if *version {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	// Pick the correct usage to call based on subcommand
	if *help {
		if rootFlags.NArg() == 0 {
			usage()
		} else {
			switch rootFlags.Arg(0) {
			case "load":
				load.Usage()
			case "list", "print":
				list.Usage()
			}
		}
		os.Exit(0)
	}

	if rootFlags.NArg() < 1 {
		usageErr(errors.New("missing subcommand"))
		os.Exit(2)
	}

	// Run subcommands
	switch rootFlags.Arg(0) {
	case "load":
		return load.Run(args, *verbose)
	case "list", "print":
		return list.Run(args, *verbose)
	case "help":
		usage()
	default:
		usageErr(fmt.Errorf("unrecognized subcommand: %s", rootFlags.Arg(0)))
		os.Exit(2)
	}
	return nil
}
