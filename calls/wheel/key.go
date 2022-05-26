package wheel

import (
	"encoding/json"
	"github.com/rjmateus/go-salt-api-client/client"
)

func keyListParser(data []byte) interface{} {
	keysData := &APIReturnType[KeysAll]{}
	json.Unmarshal(data, keysData)
	return keysData
}

func KeyListAllSync(saltClient client.SaltClient) *APIReturnType[KeysAll] {
	call := WheelCall{"key.list_all", nil, keyListParser}
	return call.CallSync(saltClient).(*APIReturnType[KeysAll])
}

func KeyFinger(match string) WheelCall {
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall{"key.finger", args, nil}
	return call
}
