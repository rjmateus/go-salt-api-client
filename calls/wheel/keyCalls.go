package wheel

import (
	"github.com/rjmateus/go-salt-api-client/client"
)

/*
* More information
* https://docs.saltproject.io/en/latest/ref/wheel/all/salt.wheel.key.html
 */

func KeyListAllSync(saltClient client.SaltClient) WheelResultType[Keys[[]string]] {
	var result = &APIReturnType[WheelResultType[Keys[[]string]]]{}
	call := WheelCall[Keys[[]string]]{"key.list_all", nil, result}
	return call.CallSync(saltClient)
}

func KeyFingerSync(saltClient client.SaltClient, match string) WheelResultType[Keys[map[string]string]] {
	var result = &APIReturnType[WheelResultType[Keys[map[string]string]]]{}

	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[Keys[map[string]string]]{"key.finger", args, result}
	return call.CallSync(saltClient)
}

func KeyAcceptSync(saltClient client.SaltClient, match string) interface{} {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.accept", args, result}
	return call.CallSync(saltClient)
}

func KeyRejectSync(saltClient client.SaltClient, match string) interface{} {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.reject", args, result}
	return call.CallSync(saltClient)
}

func KeyDeleteSync(saltClient client.SaltClient, match string) interface{} {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.delete", args, result}
	return call.CallSync(saltClient)
}

func KeyGenSync(saltClient client.SaltClient, id string) WheelResultType[KeyGen] {
	var result = &APIReturnType[WheelResultType[KeyGen]]{}
	args := make(map[string]interface{})
	args["id_"] = id
	call := WheelCall[KeyGen]{"key.gen", args, result}
	return call.CallSync(saltClient)
}

func KeyGenAcceptSync(saltClient client.SaltClient, id string, force bool) WheelResultType[KeyGen] {
	var result = &APIReturnType[WheelResultType[KeyGen]]{}
	args := make(map[string]interface{})
	args["id_"] = id
	args["force"] = force
	call := WheelCall[KeyGen]{"key.gen_accept", args, result}
	return call.CallSync(saltClient)
}
