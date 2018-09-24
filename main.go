package main

import (
	"fmt"
	"os"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

func main() {
	_, err := LookupRunfile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

func LookupRunfile() (string, error) {
	// Resolve file using the runfiles.
	file := "file.txt"
	runfile, err := bazel.Runfile(file)
	if err != nil {
		return "", err
	}
	// Check that the file actually exist
	if _, err := os.Stat(runfile); err != nil {
		return "", fmt.Errorf("File found by runfile doesn't exist")
	}	

	return "", nil
}
