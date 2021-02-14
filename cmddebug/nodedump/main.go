package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/nickatsegment/yesml"
	yaml "github.com/nickatsegment/yesml/yamlv3"
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

	// TODO support mor root types
	var out yaml.Node
	err = yesml.Unmarshal(bin, &out)
	if err != nil {
		log.Fatalf("unmarshalling: %s", err)
	}
	spew.Dump(out)

}
