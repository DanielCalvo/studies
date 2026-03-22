package main

import (
	"testing"
)

type data struct {
	name  string
	input string
	want  string
}

// table driven tests are a good fit here to test this function with multiple inputs and outputs
var testData = []data{
	{
		"in the middle of several pages",
		"<https://api.github.com/repositories/20580498/issues?after=Y3Vyc29yOnYyOpLPAAABUz0kjVjOCDz5yw%3D%3D&per_page=30&page=87>; rel=\"next\", <https://api.github.com/repositories/20580498/issues?per_page=30&page=85&before=Y3Vyc29yOnYyOpLPAAABVnYxq_DOCinWaw%3D%3D>; rel=\"prev\"",
		"https://api.github.com/repositories/20580498/issues?after=Y3Vyc29yOnYyOpLPAAABUz0kjVjOCDz5yw%3D%3D&per_page=30&page=87",
	},
	{
		"first page",
		"<https://api.github.com/repositories/20580498/issues?after=Y3Vyc29yOnYyOpLPAAABnPtyxijO86aRsA%3D%3D&per_page=30&page=2>; rel=\"next\"",
		"https://api.github.com/repositories/20580498/issues?after=Y3Vyc29yOnYyOpLPAAABnPtyxijO86aRsA%3D%3D&per_page=30&page=2",
	},
	{
		"last page",
		"<https://api.github.com/repositories/20580498/issues?per_page=30&page=87&before=Y3Vyc29yOnYyOpLPAAABUACi0dDOBnLTMw%3D%3D>; rel=\"prev\"",
		"",
	},
	{
		"pages out of order",
		"<https://api.github.com/repositories/20580498/issues?per_page=30&page=85&before=Y3Vyc29yOnYyOpLPAAABVnYxq_DOCinWaw%3D%3D>; rel=\"prev\", <https://api.github.com/repositories/20580498/issues?after=Y3Vyc29yOnYyOpLPAAABUz0kjVjOCDz5yw%3D%3D&per_page=30&page=87>; rel=\"next\"",
		"https://api.github.com/repositories/20580498/issues?after=Y3Vyc29yOnYyOpLPAAABUz0kjVjOCDz5yw%3D%3D&per_page=30&page=87",
	},
	{
		"empty header",
		"",
		"",
	},
}

func TestUrl(t *testing.T) {
	for _, tt := range testData {
		// from the docs: Run runs f as a subtest of t called name. It runs f in a separate goroutine
		// from me: Neat! It takes a test name and a test function as input
		t.Run(tt.name, func(t *testing.T) { //input is a long string, not the best experince having that as the test name on the error message
			s := getNextPageFromHeader(tt.input)
			if s != tt.want {
				t.Errorf("on test name %q, got %q, want %q", tt.name, s, tt.want)
			}
		})
	}
}
