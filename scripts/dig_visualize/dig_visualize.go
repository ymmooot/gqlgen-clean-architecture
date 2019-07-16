package main

import (
	"os"

	d "github.com/ymmooot/gqlgen-clean-architecture/app/config/dig"
	"go.uber.org/dig"
)

// Output dot file showing dependency on standard output
func main() {
	c, _ := d.BuildDigDependencies()

	dig.Visualize(c, os.Stdout)
}
