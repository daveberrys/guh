package main

import (
	"github.com/daveberrys/guh/src/cmd"
	_ "github.com/daveberrys/guh/src/cmd/create"
	_ "github.com/daveberrys/guh/src/cmd/switch"
)

func main() { cmd.Execute() }
