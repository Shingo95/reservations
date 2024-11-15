package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {

	var myH myHandler

	h := NoSurf(&myH)
	switch v := h.(type) {
	case http.Handler:
		// Do Nothing
	default:
		t.Errorf("Type is not http Handler but is %T", v)
	}
}
