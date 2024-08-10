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
			"dummy text [`test`] and [`also`] here",
			"dummy text [`test`](link#test) and [`also`](link#also) here"},
		{
			"dummy [`(*test).again`] here",
			"dummy [`(*test).again`](link#test.again) here"},
		{
			"dummy [`(*test).again`] and [`again`] here",
			"dummy [`(*test).again`](link#test.again) and [`again`](link#again) here"},
		{
			"[`(*this).one`](shouldn't change) here [`(but).this`] will",
			"[`(*this).one`](shouldn't change) here [`(but).this`](link#but.this) will"},
	}

	for i, e := range toTest {
		r := AddLinks(e.input, "link")
		if r != e.target {
			t.Errorf(`AddLinks failed! (%d)
input: %s
target: %s
return: %s`, i, e.input, e.target, r)
		}
	}
}
