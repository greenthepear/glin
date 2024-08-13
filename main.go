package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func fatalErrorCheck(err error, whileDoingWhat string) {
	if err != nil {
		log.Fatalf("Error while %s:\n\t%v", whileDoingWhat, err)
	}
}

func main() {
	in := flag.String(
		"in", "",
		"input file, default will scan from stdin")
	repoLink := flag.String(
		"repo", "",
		"link to the repository: e.g github.com/exampleguy/example")
	out := flag.String(
		"out", "",
		"output file, not setting it or -ow will just print the new text to stdout")
	overwrite := flag.Bool(
		"ow", false,
		"overwrtie the original file, will ignore -out")
	flag.Parse()

	if *repoLink == "" {
		log.Fatalf("No repo link provided, add flag like:\n\tglin -repo \"github.com/greenthepear/glin\"")
	}

	var text []byte
	var err error
	if *in == "" {
		text, err = io.ReadAll(os.Stdin)
		fatalErrorCheck(err, "reading from stdin")
	} else {
		text, err = os.ReadFile(*in)
		fatalErrorCheck(err, "reading file")
	}

	changed := AddLinks(string(text), *repoLink)
	if *overwrite {
		err = os.WriteFile(*in, []byte(changed), 0644)
		fatalErrorCheck(err, "overwriting file")
		return
	}
	if *out == "" {
		fmt.Print(changed)
		return
	}
	err = os.WriteFile(*out, []byte(changed), 0644)
	if err != nil {
		fatalErrorCheck(err, "writing to file")
	}
}
