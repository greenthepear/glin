package main

import "regexp"

func AddLinks(original []byte, repoLink string) []byte {
	// Matches [`...`] but not [`...`](, but annoyingly because of lack
	// of lookahead it matches one extra proceeding character.
	reLinkAdd := regexp.MustCompile("(\\[`.+?`\\])[^\\(]")
	reWord := regexp.MustCompile(`\w+`)

	return reLinkAdd.ReplaceAllFunc(original, func(s []byte) []byte {
		splitString := reWord.FindAll(s, -1)
		addition := []byte{}
		if len(splitString) == 1 {
			// Byte makes this ugly, probably there's a nicer way to do this
			addition = append(addition, []byte(
				"("+repoLink+"#"+string(splitString[0])+")"+string(s[len(s)-1]))...)
		} else if len(splitString) >= 2 {
			addition = append(addition, []byte(
				"("+repoLink+"#"+string(splitString[0])+"."+
					string(splitString[1])+")"+string(s[len(s)-1]))...)
		}
		return append(s[:len(s)-1], addition...)
	})
}
