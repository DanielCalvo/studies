package main

import "testing"

type data struct {
	parenthesis string
	valid       bool
}

var testData = []data{
	//{"", false},
	//{")", false},
	//{"(", false},
	//{"()", true},
	//{")(", false},
	//{"(aa)", false},
	//{"())", false},
	//{"(()", false},
	{"()()", true},
	//{"())(()", false},
	//{"(())", true},
	//{"(()(()))", true},
}

func TestAreParenthesisValid(t *testing.T) {
	for _, td := range testData {
		t.Run("Parenthesis: "+td.parenthesis, func(t *testing.T) {
			got := AreParenthesisValid(td.parenthesis)
			want := td.valid
			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}
