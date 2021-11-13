package urlparser

import "testing"

func TestParser(t *testing.T) {
	urls := []string{
		"https://pkg.go.dev/golang.org/x/text/unicode/rangetable",
		"https://pkg.go.dev/golang.org/x/text/unicode/rangetable#New",
		"https://www.includehelp.com/golang/unicode-isoneof-function-with-examples.aspx",
		"https://sso.a1.bg/oauth/2/login",
	}

	for _, v := range urls {
		if err := Parse(v); err != nil {
			t.Fatal(err)
		}
	}
}
