package wheel

import (
	"github.com/rjmateus/go-salt-api-client/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestKeyListAllSync(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 1, "name": "kyle", "description": "novice gopher"}`))
	}))
	client.NewClient(httpServer.URL, "", "", "")

}
