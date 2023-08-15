package main

import "testing"

type data struct {
	s                    string
	c                    string
	doesithavechar3times bool
}

var testData = []data{
	{"banana", "a", true},
	{"cccombobreaker", "c", true},
	{"how about no", "b", false},
	{"tooooo many o's, yo", "o", false},
	//{"asdasdsadsadsaa", "a", true}, //fails as expected
}

func TestDoesItHaveChar3Times(t *testing.T) {
	for _, td := range testData {
		//You can define other name in here, but I'm using the string for a name
		t.Run(td.s, func(t *testing.T) {
			got := DoesItHaveChar3Times(td.s, td.c)
			want := td.doesithavechar3times
			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}
