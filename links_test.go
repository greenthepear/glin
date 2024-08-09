package main

import (
	"testing"
)

func TestAddLinks(t *testing.T) {
	type inputAndExpected struct {
		input, target string
	}

	toTest := []inputAndExpected{
		{
			"dummy text [`test`] here",
			"dummy text [`test`](link#test) here"},
		{
			"dummy [`(*test).again`] here",
			"dummy [`(*test).again`](link#test.again) here"},
		{
			"dummy [`(*test).again`] here [`again`] ",
			"dummy [`(*test).again`](link#test.again) here [`again`](link#again) "},
	}

	for i, e := range toTest {
		r := string(AddLinks([]byte(e.input), "link"))
		if r != e.target {
			t.Errorf(`AddLinks failed! (%d)
input: %s
target: %s
return: %s`, i, e.input, e.target, r)
		}
	}
}
