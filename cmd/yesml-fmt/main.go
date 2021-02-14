package main

import (
	"io/ioutil"
	"log"
	"os"

	yesmlfmt "github.com/nickatsegment/yesml/fmt"
)

func main() {
	fn := "-"
	if len(os.Args) > 1 {
		fn = os.Args[1]
	}
	var bin []byte
	var err error

	if fn == "-" {
		bin, err = ioutil.ReadAll(os.Stdin)
	} else {
		bin, err = ioutil.ReadFile(fn)
	}
	if err != nil {
		log.Fatalf("reading from %q: %s", fn, err)
	}

	bout, err := yesmlfmt.Format(bin)

	if err != nil {
		log.Fatalf("formatting: %s", err)
	}
	os.Stdout.Write(bout)
}
