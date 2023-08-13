package main

import (
	"flag"
	"os"

	"github.com/shimabukuromeg/mycat"
)

var n bool

func init() {
	flag.BoolVar(&n, "n", false, "number all output lines mode")
}

func main() {
	flag.Parse()
	args := flag.Args()
	mycat.Read(os.Stdout, args, n)
}
