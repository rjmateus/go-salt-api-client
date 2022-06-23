package wheel

import (
	"github.com/rjmateus/go-salt-api-client/testUtils"
	"reflect"
	"testing"
)

func TestKeyListAllSyncOK(t *testing.T) {

	call := KeyListAll()
	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun": "key.list_all",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}
	data := call.CallSync(htppClient)
	if len(data.Data.Return.Local) != 2 || len(data.Data.Return.MinionsPre) != 1 || len(data.Data.Return.MinionsRejected) != 1 || len(data.Data.Return.MinionsRejected) != 1 {
		t.Errorf("data not correctly loaded and de-serialized")
	}
}

func TestKeyFingerSyncOK(t *testing.T) {
	call := KeyFinger("my_id")
	//	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun":   "key.finger",
		"match": "my_id",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}
	/*data := call.CallSync(htppClient)
	if len(data.Data.Return.Local) != 2 || len(data.Data.Return.MinionsPre) != 1 || len(data.Data.Return.MinionsRejected) != 1 || len(data.Data.Return.MinionsRejected) != 1 {
		t.Errorf("data not correctly loaded and de-serialized")
	}*/
}

func TestKeyAcceptSyncOK(t *testing.T) {

	call := KeyAccept("my_id")
	//	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun":   "key.accept",
		"match": "my_id",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}

	/*data := call.CallSync(htppClient)
	if len(data.Data.Return.Local) != 2 || len(data.Data.Return.MinionsPre) != 1 || len(data.Data.Return.MinionsRejected) != 1 || len(data.Data.Return.MinionsRejected) != 1 {
		t.Errorf("data not correctly loaded and de-serialized")
	}*/
}

func TestKeyRejectSyncOK(t *testing.T) {
	call := KeyReject("my_id")
	//	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun":   "key.reject",
		"match": "my_id",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}

	/*data := call.CallSync(htppClient)
	if len(data.Data.Return.Local) != 2 || len(data.Data.Return.MinionsPre) != 1 || len(data.Data.Return.MinionsRejected) != 1 || len(data.Data.Return.MinionsRejected) != 1 {
		t.Errorf("data not correctly loaded and de-serialized")
	}*/
}

func TestKeyDeleteSyncOK(t *testing.T) {
	call := KeyDelete("my_id")
	//	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun":   "key.delete",
		"match": "my_id",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}
	/*data := call.CallSync(htppClient)
	if len(data.Data.Return.Local) != 2 || len(data.Data.Return.MinionsPre) != 1 || len(data.Data.Return.MinionsRejected) != 1 || len(data.Data.Return.MinionsRejected) != 1 {
		t.Errorf("data not correctly loaded and de-serialized")
	}*/
}

func TestKeyGenSyncOK(t *testing.T) {
	call := KeyGen("my_id")
	//	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)

	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun": "key.gen",
		"id_": "my_id",
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)
	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}
	/*data := call.CallSync(htppClient)
	if len(data.Data.Return.Local) != 2 || len(data.Data.Return.MinionsPre) != 1 || len(data.Data.Return.MinionsRejected) != 1 || len(data.Data.Return.MinionsRejected) != 1 {
		t.Errorf("data not correctly loaded and de-serialized")
	}*/
}

func TestKeyGenAcceptSyncOK(t *testing.T) {
	call := KeyGenAccept("my_id", true)
	//	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)

	if call.Fun != "key.gen_accept" {
		t.Errorf("wrong fun name. got: %s, want: %s", call.Fun, "key.gen_accept")
	}
	payload := call.GetPayload()
	expectedPayload := map[string]interface{}{
		"fun":   "key.gen_accept",
		"id_":   "my_id",
		"force": true,
	}
	checkPayload := reflect.DeepEqual(payload, expectedPayload)

	if !checkPayload {
		t.Errorf("Payload not expected. Got %s -> expects: %s", payload, expectedPayload)
	}

	/*data := call.CallSync(htppClient)
	if len(data.Data.Return.Local) != 2 || len(data.Data.Return.MinionsPre) != 1 || len(data.Data.Return.MinionsRejected) != 1 || len(data.Data.Return.MinionsRejected) != 1 {
		t.Errorf("data not correctly loaded and de-serialized")
	}*/
}
