package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/inconshreveable/log15"
)

var (
	token = "asdoim896543OHFdhcluyhsdh" //made up token
	text  = "Hello World!"
)

var cowsayPath string

func init() {
	flag.StringVar(&cowsayPath, "cowsayPath", "/usr/local/bin/cowsay", "path to cowsay executable")
	flag.Parse()
	fmt.Printf("INIT flag=%v", cowsayPath)
}

func TestPretty(t *testing.T) {
	log := log15.New()
	h := cowsayHandler(
		cowsayPath,
		nil,
		log,
	)
	r, w := getRW(t)
	h(w, r)
	expected := 263
	if len(w.b) != expected {
		t.Errorf("Incorrect response length : want %d, got %d", expected, len(w.b))
	}
	if w.s != http.StatusOK {
		t.Errorf(
			"Wrong status code: Expected %d, Got %d",
			http.StatusOK,
			w.s,
		)
	}

}
func TestMissingProg(t *testing.T) {
	log := log15.New()
	h := cowsayHandler(
		"/bad/path",
		nil,
		log,
	)

	r, w := getRW(t)
	h(w, r)
	if len(w.b) > 0 {
		t.Error("Expected no response from bad cowsay path")
	}
	if w.s != http.StatusInternalServerError {
		t.Errorf(
			"Wrong status code: Expected %d, Got %d",
			http.StatusInternalServerError,
			w.s,
		)
	}
}

func TestWrongToken(t *testing.T) {
	log := log15.New()
	tokens := make(map[string]bool)
	tokens["bad-token"] = true
	h := cowsayHandler(
		cowsayPath,
		tokens,
		log,
	)

	r, w := getRW(t)
	h(w, r)
	if len(w.b) > 0 {
		t.Error("Expected no response from wrong token")
	}
	if w.s != http.StatusForbidden {
		t.Errorf(
			"Wrong status code: Expected %d, Got %d",
			http.StatusForbidden,
			w.s,
		)
	}

}

func TestCorrectToken(t *testing.T) {
	log := log15.New()
	tokens := make(map[string]bool)
	tokens[token] = true
	h := cowsayHandler(
		cowsayPath,
		tokens,
		log,
	)

	r, w := getRW(t)
	h(w, r)
	expected := 263
	if len(w.b) != expected {
		t.Errorf("Incorrect response length : want %d, got %d", expected, len(w.b))
	}
	if w.s != http.StatusOK {
		t.Errorf(
			"Wrong status code: Expected %d, Got %d",
			http.StatusOK,
			w.s,
		)
	}

}
func getRW(t *testing.T) (*http.Request, *mockWriter) {
	return getTestRequest(t), getTestWriter(t)
}

func getTestRequest(t *testing.T) *http.Request {
	data := url.Values{}
	data.Set("token", token)
	data.Set("text", text)
	r, err := http.NewRequest("POST", "", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.Error(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	return r

}
func getTestWriter(t *testing.T) *mockWriter {
	return &mockWriter{h: make(map[string][]string)}
}

type mockWriter struct {
	h http.Header
	s int
	b []byte
}

func (w *mockWriter) Header() http.Header {
	return w.h
}
func (w *mockWriter) Write(b []byte) (int, error) {
	w.b = b
	return len(b), nil
}
func (w *mockWriter) WriteHeader(i int) {
	w.s = i

}
