package load

import (
	"fmt"
	"os"

	"github.com/bensallen/modscan/pkg/mod"
	"github.com/bensallen/modscan/pkg/uname"
	flag "github.com/spf13/pflag"
	"github.com/u-root/u-root/pkg/kmodule"
)

const usageHeader = `load - Load needed kernel modules

Usage:
  load

Flags:
`

var (
	flags    = flag.NewFlagSet("load", flag.ContinueOnError)
	rootpath = flags.StringP("root", "r", "/", "Root path for kernel modules")
	kver     = flags.StringP("kver", "k", "", "Set kernel version instead of using uname")
	syspath  = flags.StringP("syspath", "s", "/sys/devices", "Path to sysfs devices")
)

// Usage of the load subcommand
func Usage() {
	fmt.Fprintf(os.Stderr, usageHeader)

	// Set the modpath flag's default to use uname release.
	u, _ := uname.New()
	flags.Lookup("kver").DefValue = u.Release()

	fmt.Fprintf(os.Stderr, flags.FlagUsagesWrapped(0)+"\n")
}

// Run load subcommand
func Run(args []string, verbose bool) error {
	// Set the modpath flag's default to use uname release.
	u, err := uname.New()
	flags.Lookup("kver").Value.Set(u.Release())

	flags.ParseErrorsWhitelist.UnknownFlags = true
	if err := flags.Parse(args); err != nil {
		Usage()
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		os.Exit(2)
	}

	modpath := mod.ModuleDir(*rootpath, *kver)

	modAliases, err := mod.ParseKernelModulesAlias(modpath)
	if err != nil {
		return err
	}

	modules, err := mod.WalkModAlias(*syspath, modAliases, nil)
	if err != nil {
		return err
	}

	kopts := kmodule.ProbeOpts{RootDir: *rootpath, KVer: *kver}

	for modName := range modules {
		if verbose {
			fmt.Printf("loading module %q\n", modName)
		}
		if err := kmodule.ProbeOptions(modName, "", kopts); err != nil {
			fmt.Printf("could not load module %q: %v\n", modName, err)
		}
	}

	return nil
}
