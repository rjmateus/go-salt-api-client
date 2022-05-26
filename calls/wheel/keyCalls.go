package wheel

import (
	"github.com/rjmateus/go-salt-api-client/client"
)

/*
* More information
* https://docs.saltproject.io/en/latest/ref/wheel/all/salt.wheel.key.html
 */

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

func KeyRejectSync(saltClient client.SaltClient, match string) interface{} {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.reject", args, result}
	call.CallSync(saltClient)
	return call.result.Return[0].Data.Return
}

func KeyDeleteSync(saltClient client.SaltClient, match string) interface{} {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.delete", args, result}
	call.CallSync(saltClient)
	return call.result.Return[0].Data.Return
}

func KeyGenSync(saltClient client.SaltClient, id string) KeyGen {
	var result = &APIReturnType[WheelResultType[KeyGen]]{}
	args := make(map[string]interface{})
	args["id_"] = id
	call := WheelCall[KeyGen]{"key.gen", args, result}
	call.CallSync(saltClient)
	return call.result.Return[0].Data.Return
}

func KeyGenAcceptSync(saltClient client.SaltClient, id string, force bool) KeyGen {
	var result = &APIReturnType[WheelResultType[KeyGen]]{}
	args := make(map[string]interface{})
	args["id_"] = id
	args["force"] = force
	call := WheelCall[KeyGen]{"key.gen_accept", args, result}
	call.CallSync(saltClient)
	return call.result.Return[0].Data.Return
}
