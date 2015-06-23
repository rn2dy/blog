package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRouterPathMatcher(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{in: "/", out: "/$"},
		{in: "/articles", out: "/articles$"},
		{in: "/articles/:slug", out: "/articles/([^/?]+)$"},
		{in: "/articles/:slug/:part", out: "/articles/([^/?]+)/([^/?]+)$"},
		{in: "/articles/:slug/2015/:part/", out: "/articles/([^/?]+)/2015/([^/?]+)$"},
	}

	r := &Router{}
	for i, c := range cases {
		r.add(c.in, nil)
		if len(r.matchers) != i+1 {
			t.Errorf("Number of matchers is %d, want %d", len(r.matchers), i+1)
		}
		m := r.matchers[i]
		if m.reg.String() != c.out {
			t.Errorf("Matcher regex string is %q, want %q", m.reg.String(), c.out)
		}
	}
}

func TestRouterPathVariableNames(t *testing.T) {
	cases := []struct {
		in  string
		out []string
	}{
		{in: "/articles", out: []string{}},
		{in: "/articles/:slug", out: []string{"slug"}},
		{in: "/articles/:slug/:part", out: []string{"slug", "part"}},
		{in: "/articles/:slug/2015/:part/", out: []string{"slug", "part"}},
	}

	r := &Router{}

	for i, c := range cases {
		r.add(c.in, nil)
		m := r.matchers[i]
		if len(m.names) != len(c.out) {
			t.Errorf("Number of matcher variable names %d, want %d", len(m.names), len(c.out))
		} else {
			for i := 0; i < len(m.names); i++ {
				if m.names[i] != c.out[i] {
					t.Errorf("Variable names is %q, want %q", m.names[i], c.out[i])
				}
			}
		}
	}
}

func TestRouteMatching(t *testing.T) {
	cases := []struct {
		in1, in2 string
		out      map[string]string
	}{
		{
			in1: "/articles/:slug",
			in2: "http://localhost:3000/articles/tutorial",
			out: map[string]string{
				"slug": "tutorial",
			},
		},
		{
			in1: "/articles/:slug/:part",
			in2: "http://localhost:3000/articles/tutorial/one",
			out: map[string]string{
				"slug": "tutorial",
				"part": "one",
			},
		},
	}

	r := &Router{}
	for _, c := range cases {
		r.add(c.in1, nil)

		ctx := &C{vars: make(map[string]string)}
		req, _ := http.NewRequest("GET", c.in2, nil)
		_, err := r.match(ctx, req)
		if err != nil {
			t.Error(err)
		}
		for k, _ := range c.out {
			if _, ok := ctx.vars[k]; !ok {
				t.Errorf("Value of %q not extracted from %q", k, c.in2)
			}
		}
	}
}

func handlerWare(ch chan string, t *testing.T, cases [][]string) func(*C, http.ResponseWriter, *http.Request) {
	return func(ctx *C, w http.ResponseWriter, req *http.Request) {
		ch <- req.URL.Path
		for _, c := range cases {
			if ctx.vars[c[0]] != c[1] {
				t.Errorf("Path variable has value %q, but want %q", ctx.vars[c[0]], c[1])
			}
		}
	}
}

func TestRouting(t *testing.T) {
	// config router
	var ch = make(chan string, 2)
	r := &Router{}
	r.add("/articles/:slug", handlerWare(ch, t, [][]string{{"slug", "tutorial"}}))
	r.add("/articles/:slug/:part", handlerWare(ch, t, [][]string{{"slug", "tutorial", "part", "one"}}))

	// send request
	var urls = []string{
		"http://localhost:3000/articles/tutorial",
		"http://localhost:3000/articles/tutorial/one",
	}
	go func() {
		for _, url := range urls {
			ctx := &C{vars: make(map[string]string)}
			req, _ := http.NewRequest("GET", url, nil)
			r.route(ctx, httptest.NewRecorder(), req)
		}
	}()

	// verify
	var reqcount = 0
	var paths []string
	for {
		select {
		case path := <-ch:
			reqcount++
			for _, p := range paths {
				if p == path {
					t.Errorf("Path %q gets routed more than once.", path)
				}
			}
			paths = append(paths, path)
			if reqcount == 2 {
				return
			}
		case <-time.After(5 * time.Second):
			t.Fatal("Http request timing out.")
		}
	}
}
