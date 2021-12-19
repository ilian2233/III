package html_parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParser(t *testing.T) {
	content, err := ioutil.ReadFile("./test.html")     // the file is inside the local directory
	if err != nil {
		t.Error(err)
	}

	symbols, err := html_document(string(content))
	if err != nil {
		t.Error(err)
	}
	
	fmt.Println(symbols)
}
