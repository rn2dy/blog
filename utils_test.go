package main

import (
	"reflect"
	"testing"
)

func TestArticle(t *testing.T) {
	cases := []struct {
		in  string
		out *Article
	}{
		{
			in: "2015-01-02_ruby-meta-programming-tips.md",
			out: &Article{
				Title: "meta programming tips",
				Tags:  []string{"ruby"},
				Date:  makeDate(2015, 1, 2),
				src:   "2015-01-02_ruby-meta-programming-tips.md",
			},
		},
		{
			in: "2014-05-02_golang_ruby-function-adapters.md",
			out: &Article{
				Title: "function adapters",
				Tags:  []string{"golang", "ruby"},
				Date:  makeDate(2014, 5, 2),
				src:   "2014-05-02_golang_ruby-function-adapters.md",
			},
		},
		{
			in: "2011-12-08_javascript-prototype-inheritance.md",
			out: &Article{
				Title: "prototype inheritance",
				Tags:  []string{"javascript"},
				Date:  makeDate(2011, 12, 8),
				src:   "2011-12-08_javascript-prototype-inheritance.md",
			},
		},
	}

	for _, c := range cases {
		a := NewArticle(c.in)

		if !reflect.DeepEqual(a, c.out) {
			t.Errorf("%q parsed as %v, but want %v", c.in, a, c.out)
		}
	}
}
