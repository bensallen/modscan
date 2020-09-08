package list

import (
	"fmt"
	"os"

	"github.com/bensallen/modscan/pkg/mod"
	"github.com/bensallen/modscan/pkg/uname"
	flag "github.com/spf13/pflag"
)

const usageHeader = `list - Scan for needed kernel modules and list the results

Usage:
  [list|print]

Flags:
`

var (
	flags   = flag.NewFlagSet("list", flag.ContinueOnError)
	modpath = flags.StringP("modpath", "m", "", "Path to kernel modules")
	syspath = flags.StringP("syspath", "s", "/sys/devices", "Path to sysfs devices")
)

// Usage of the list subcommand
func Usage() {
	fmt.Fprintf(os.Stderr, usageHeader)

	// Set the modpath flag's default to use uname release.
	u, _ := uname.New()
	flags.Lookup("modpath").DefValue = "/lib/modules/" + u.Release()

	fmt.Fprintf(os.Stderr, flags.FlagUsagesWrapped(0)+"\n")
}

// Run list subcommand
func Run(args []string, verbose bool) error {
	// Set the modpath flag's default to use uname release.
	u, err := uname.New()
	flags.Lookup("modpath").Value.Set("/lib/modules/" + u.Release())

	flags.ParseErrorsWhitelist.UnknownFlags = true
	if err := flags.Parse(args); err != nil {
		Usage()
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		os.Exit(2)
	}

	modAliases, err := mod.ParseKernelModulesAlias(*modpath)
	if err != nil {
		return err
	}

	_, err = mod.WalkModAlias(*syspath, modAliases, os.Stdout)
	if err != nil {
		return err
	}

	return nil
}
