package main

import (
	"github.com/belendia/core"
)

func main() {
	a := core.App{}
	a.Port = ":9003"
	a.Initialize()
	a.Run()
}
