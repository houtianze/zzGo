package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var appname string

func usage() {
	var format = `Check if the given path is already in the PATH environment variable.\n
Exit code:
  0 if in
  1 if NOT in
  2 on usage error
Usage:
  %s <path>
`
	fmt.Fprintf(os.Stderr, format, appname)
	os.Exit(2)
}

func samepath(a string, b string) bool {
	var aa = filepath.Clean(a)
	var bb = filepath.Clean(b)

	if runtime.GOOS == "windows" {
		return strings.ToUpper(aa) == strings.ToUpper(bb)
	} else {
		return aa == bb
	}
}

func isin(given string) bool {
	// TODO: What if PathListSeparator is inside each path? Nee a better PATH string parser
	// but live with the simple splitting for now
	var paths = strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	for _, e := range paths {
		if samepath(e, given) {
			return true
		}
	}

	return false
}

func main() {
	appname = filepath.Base(os.Args[0])
	if len(os.Args) != 2 {
		usage()
	} else {
		if !isin(os.Args[1]) {
			os.Exit(1)
		}
	}
}
