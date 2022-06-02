package testUtils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func GetHttpServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		bytes, error := ioutil.ReadFile("all_keys_test.json")
		if error != nil {
			log.Fatal(error)
		}
		w.Write(bytes)
	}))

}
