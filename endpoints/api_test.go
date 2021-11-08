package endpoints

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestShortUrl(t *testing.T) {

	router := httprouter.New()
	router.POST("/", ShortUrl)

	req, _ := http.NewRequest("POST", "localhost:8081/?url=https://www.google.com", nil)
	rr := httptest.NewRecorder()
	fmt.Println(rr.Code)

	router.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Fatal("Test case passed")
	}
}