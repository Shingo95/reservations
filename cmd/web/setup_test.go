package main

import (
	"net/http"
	"os"
	"testing"
)

type myHandler struct{}

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
func (myH *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
