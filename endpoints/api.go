package endpoints

import (
	"fmt"
	"go-tiny-url/services"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var root_url string = "http://infra.io/"

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func ShortUrl(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	code := req.URL.Query().Get("url")
	shorturl := services.UrlsDataSet(code)
	fmt.Fprintf(w, "%s\n", root_url+shorturl)
}