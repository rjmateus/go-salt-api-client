package wheel

import (
	"github.com/rjmateus/go-salt-api-client/testUtils"
	"testing"
)

func TestKeyListAllSyncOK(t *testing.T) {
	call := KeyListAllSync()
	htppClient := testUtils.MockedSaltFunResponseClient(call.Fun)

	if call.Fun != "key.list_all" {
		t.Errorf("wrong fun name. got: %s, want: 5%s", call.Fun, "key.list_all")
	}
	data := call.CallSync(htppClient)
	if len(data.Data.Return.Local) != 2 || len(data.Data.Return.MinionsPre) != 1 || len(data.Data.Return.MinionsRejected) != 1 || len(data.Data.Return.MinionsRejected) != 1 {
		t.Errorf("data not de-serialized")
	}
}
