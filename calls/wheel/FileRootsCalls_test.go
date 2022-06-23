package wheel

import (
	"github.com/rjmateus/go-salt-api-client/testUtils"
	"reflect"
	"testing"
)

func TestFileRootsListRootSyncOK(t *testing.T) {

	call := FileRootsListRoot()
	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun": "file_roots.list_roots",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}
	data := call.CallSync(htppClient)
	if len(data.Data.Return.Base) != 1 || len(data.Data.Return.Base[0]) != 2 {
		t.Errorf("data not correctly loaded and de-serialized")
	}
}

func TestFileRootsListEnvCallEmptyEnv(t *testing.T) {

	call := FileRootsListEnv("")
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun": "file_roots.list_env",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}
}

func TestFileRootsListEnvCallEnv(t *testing.T) {

	call := FileRootsListEnv("my_env")
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun":     "file_roots.list_env",
		"saltenv": "my_env",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}
}
