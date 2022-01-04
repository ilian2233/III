package http_client

import (
	"testing"
)

func TestVisitURL(t *testing.T) {

	//"https://ednsquare.com/story/how-to-make-http-requests-in-golang------5VIjL3"

	if err := VisitURL(); err != nil {
		t.Fatal(err)
	}
}
