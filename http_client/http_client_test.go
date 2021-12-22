package http_client

import (
	"testing"
)

func TestVisitURL(t *testing.T) {

	if err := VisitURL("https://ednsquare.com/story/how-to-make-http-requests-in-golang------5VIjL3"); err != nil {
		t.Fatal(err)
	}
}
