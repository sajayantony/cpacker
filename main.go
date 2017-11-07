package main

import (
	"flag"
	"fmt"

	"github.com/sajayantony/cpacker/packer"
)

var packDirectory string

func init() {
	flag.StringVar(&packDirectory, "dir", ".", "directory to pack")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, world.\n")
	packer.EnsureDockerfile()
	packer.Pack()
}
