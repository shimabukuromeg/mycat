package main

import (
	"flag"

	"github.com/shimabukuromeg/mycat"
)

var n bool

func init() {
	flag.BoolVar(&n, "n", false, "number all output lines mode")
}

func main() {
	flag.Parse()
	args := flag.Args()
	mycat.Read(args, n)
}
