package main

import (
	"regexp"
)

func AddLinks(original string, repoLink string) string {
	// \[`[^\[\]]+?`\][^(]
	// Matches [`...`] but not [`...`](, but annoyingly because of lack
	// of lookahead it matches one extra proceeding character.
	reLinkAdd := regexp.MustCompile(`\[` + "`" + `[^\[\]]+?` + "`" + `\][^(]`)
	reWord := regexp.MustCompile(`\w+`)
	return reLinkAdd.ReplaceAllStringFunc(original, func(s string) string {
		splitString := reWord.FindAllString(s, -1)
		addition := ""
		if len(splitString) == 1 {
			addition =
				"(" + repoLink + "#" + string(splitString[0]) + ")" + string(s[len(s)-1])
		} else if len(splitString) >= 2 {
			addition =
				"(" + repoLink + "#" + string(splitString[0]) + "." +
					string(splitString[1]) + ")" + string(s[len(s)-1])
		}
		return s[:len(s)-1] + addition
	})
}
