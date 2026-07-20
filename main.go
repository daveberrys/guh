package main

import (
	"github.com/daveberrys/guh/src/cmd"
	_ "github.com/daveberrys/guh/src/cmd/account"
	_ "github.com/daveberrys/guh/src/cmd/branch"
	_ "github.com/daveberrys/guh/src/cmd/browse"
	_ "github.com/daveberrys/guh/src/cmd/commit"
	_ "github.com/daveberrys/guh/src/cmd/diff"
	_ "github.com/daveberrys/guh/src/cmd/init"
	_ "github.com/daveberrys/guh/src/cmd/link"
	_ "github.com/daveberrys/guh/src/cmd/logs"
	_ "github.com/daveberrys/guh/src/cmd/pull"
	_ "github.com/daveberrys/guh/src/cmd/push"
	_ "github.com/daveberrys/guh/src/cmd/stash"
	_ "github.com/daveberrys/guh/src/cmd/undo"
	_ "github.com/daveberrys/guh/src/cmd/cli"
)

func main() { cmd.Execute() }
