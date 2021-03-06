package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHello(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(hello))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}
	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}
