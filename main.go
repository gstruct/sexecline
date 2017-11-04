package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	path := os.Getenv("PATH")
	if len(path) > 0 {
		path = fmt.Sprintf("%s:%s", dir, path)
	} else {
		path = dir
	}
	os.Setenv("PATH", path)
	err = syscall.Exec(filepath.Join(dir, "execlineb"), os.Args, syscall.Environ())
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
