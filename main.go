package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	in := flag.String("in", "", "input markdown file")
	repoLink := flag.String("repo", "", "link to the repository: e.g github.com/exampleguy/example")
	out := flag.String(
		"out", "",
		"output file, not setting it will ask to confirm overwriting, use -o to not ask")
	overwrite := flag.Bool(
		"ow", false,
		"overwrtie the original file, will ignore -out and not bug you to confirm overwrite if output no provided")

	flag.Parse()
	if *in == "" {
		log.Fatalf("No input file provided.")
	}
	if *repoLink == "" {
		log.Fatalf("No repo link provided, add flag like:\n\tlinktime -repo \"github.com/greenthepear/linktime\"")
	}

	text, err := os.ReadFile(*in)
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	changed := AddLinks(string(text), *repoLink)
	if *out == "" {
		if !*overwrite {
			fmt.Printf("Do you want to overwrite %s? (y/N)\n", *in)
			choice := "N"
			_, err := fmt.Scan(&choice)
			if err != nil {
				log.Fatalf("Bad input.")
			}
			if choice != "Y" && choice != "y" {
				return
			}
		}
		err := os.WriteFile(*in, []byte(changed), 0644)
		if err != nil {
			log.Fatalf("Error while writing file: %v", err)
		}
	}

}
