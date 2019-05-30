package main

import (
	"github.com/lvchengchang/cdb/cdb"
)

func main() {
	op := cdb.NewOptions()
	op.Logger.LogDebug("option success")
}
