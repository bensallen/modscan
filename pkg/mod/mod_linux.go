package mod

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fanyang01/radix"
)

// Alias is a pattern and matching module parsed from a modules.alias file.
type Alias struct {
	pattern string
	module  string
}

// ParseKernelModulesAlias loads the patterns and the resulting module name
// found in a modules.alias file into a radix.PatternTrie for efficient
// searching. Supplied path should be the directory that contains a modules.alias
// file. Parsing modules.alias.bin not supported.
func ParseKernelModulesAlias(path string) (*radix.PatternTrie, error) {
	f, err := os.Open(path + "/modules.alias")
	defer f.Close()
	if err != nil {
		return nil, err
	}

	modAliases := radix.NewPatternTrie()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		alias := strings.Split(txt, " ")
		if len(alias) == 3 {
			modAlias, modName := alias[1], alias[2]
			modAliases.Add(modAlias, Alias{modAlias, modName})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return modAliases, nil
}

// WalkModAlias walks through sysfs looking for modalias files, reading the alias,
// and searching for a matching module name in the provided radix.PatternTrie.
// Wildcards (*) are supported in the provided patterns in radix.PatternTrie.
func WalkModAlias(path string, modAliases *radix.PatternTrie, output *os.File) (map[string]bool, error) {
	modules := map[string]bool{}

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Base(path) == "modalias" {
				f, err := os.Open(path)
				defer f.Close()
				if err != nil {
					return err
				}
				aliasByte, err := ioutil.ReadAll(f)
				if err != nil {
					return err
				}
				alias := strings.TrimSpace(string(aliasByte))

				if moduleName, ok := modAliases.Lookup(alias); ok {
					if output != nil {
						fmt.Fprintf(output, "Device %s with alias %s needs module %s, matched on pattern %s\n", filepath.Dir(path), alias, moduleName.(Alias).module, moduleName.(Alias).pattern)
					}
					modules[moduleName.(Alias).module] = true
				}
			}
			return nil
		})

	return modules, err
}

// ModuleDir attempts to find a kernel module directory based on the provided
// root and release arguments. Returns the path if found or an empty string
// if not found.
func ModuleDir(root string, release string) string {
	var moduleDir string
	for _, path := range []string{"/lib/modules", "/usr/lib/modules"} {
		moduleDir = filepath.Join(root, path, release)
		if _, err := os.Stat(moduleDir); err == nil {
			break
		}
	}
	return moduleDir
}
