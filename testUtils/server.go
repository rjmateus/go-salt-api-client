package testUtils

import (
	"fmt"
	"github.com/rjmateus/go-salt-api-client/client"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path"
	"runtime"
)

func getHttpServerResponseFile(fileName string, status int) *httptest.Server {
	_, thisFileName, _, _ := runtime.Caller(1)

	bytes, error := ioutil.ReadFile(path.Join(path.Dir(thisFileName), fileName))
	if error != nil {
		log.Fatal(error)
	}

	return GetHttpServer(bytes, status)
}

func GetHttpServer(response []byte, status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(response)
	}))
}

func MockedSaltFunResponseClient(funName string) client.SaltClient {
	http := getHttpServerResponseFile(fmt.Sprintf("%s_reponse.json", funName), http.StatusOK)
	htppClient := client.NewClient(http.URL, "", "", "")
	return htppClient
}
