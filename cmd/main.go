package main

import (
	"flag"
	"github.com/lvchengchang/cdb/cdb"
	"log"
	"os"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")

	flag.Parse()
}

func main() {
	op, err := cdb.NewOptions(confPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	op.Logger.RecordLogData("option success")
	c, err := cdb.Open("cdb")
	if err != nil {
		log.Panicln(err)
	}
	defer c.Close()
}
