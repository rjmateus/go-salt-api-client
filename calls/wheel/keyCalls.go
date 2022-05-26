package wheel

import (
	"github.com/rjmateus/go-salt-api-client/client"
)

func KeyListAllSync(saltClient client.SaltClient) Keys[[]string] {
	var result = &APIReturnType[WheelResultType[Keys[[]string]]]{}
	call := WheelCall[Keys[[]string]]{"key.list_all", nil, result}
	call.CallSync(saltClient)
	return call.result.Return[0].Data.Return
}

func KeyFingerSync(saltClient client.SaltClient, match string) Keys[map[string]string] {
	var result = &APIReturnType[WheelResultType[Keys[map[string]string]]]{}

	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[Keys[map[string]string]]{"key.finger", args, result}
	call.CallSync(saltClient)
	return call.result.Return[0].Data.Return
}

func KeyAcceptSync(saltClient client.SaltClient, match string) interface{} {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.accept", args, result}
	call.CallSync(saltClient)
	return call.result.Return[0].Data.Return
}
