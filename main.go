package main

import (
	"bytes"
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

func getRepoFromGoMod() (string, error) {
	full, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	full = bytes.SplitN(full, []byte("\n"), 2)[0]
	return string(bytes.SplitN(full, []byte(" "), 3)[1]), nil
}

func main() {
	in := flag.String(
		"in", "",
		"input file, omitting will scan from stdin")
	repoLink := flag.String(
		"repo", "",
		"path to the repository, e.g. github.com/exampleguy/example\nomitting will try to get the path from go.mod in the working directory")
	out := flag.String(
		"out", "",
		"output file, omitting will print the new text to stdout")
	overwrite := flag.Bool(
		"ow", false,
		"overwrite the original file, will ignore -out")
	flag.Parse()

	var err error
	if *repoLink == "" {
		gomodRepo, err := getRepoFromGoMod()
		if err != nil || gomodRepo == "" {
			log.Fatalf("No repo path provided or found in go.mod, add flag like:\n\tglin -repo \"github.com/greenthepear/glin\"\nError while opening go.mod (if any): %v", err)
		}
		repoLink = &gomodRepo
	}

	var text []byte
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
