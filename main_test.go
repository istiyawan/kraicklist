package main

import (
	// "fmt"
	// "github.com/stretchr/testify/assert"
	// "log"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleSearch(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/search?q=iphone", nil)
	w := httptest.NewRecorder()

	// handleSearch()

	searcher := &Searcher{}
	err := searcher.Load("data.gz")
	if err != nil {
		log.Fatalf("unable to load search data due: %v", err)
	}

	handleSearch(searcher).ServeHTTP(w, r)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}

}
