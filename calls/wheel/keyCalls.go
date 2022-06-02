package wheel

/*
* More information
* https://docs.saltproject.io/en/latest/ref/wheel/all/salt.wheel.key.html
 */

func KeyListAllSync() WheelCall[Keys[[]string]] {
	var result = &APIReturnType[WheelResultType[Keys[[]string]]]{}
	call := WheelCall[Keys[[]string]]{"key.list_all", nil, result}
	return call
}

func KeyFingerSync(match string) WheelCall[Keys[map[string]string]] {
	var result = &APIReturnType[WheelResultType[Keys[map[string]string]]]{}

	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[Keys[map[string]string]]{"key.finger", args, result}
	return call
}

func KeyAcceptSync(match string) WheelCall[interface{}] {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.accept", args, result}
	return call
}

func KeyRejectSync(match string) WheelCall[interface{}] {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.reject", args, result}
	return call
}

func KeyDeleteSync(match string) WheelCall[interface{}] {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.delete", args, result}
	return call
}

func KeyGenSync(id string) WheelCall[KeyGen] {
	var result = &APIReturnType[WheelResultType[KeyGen]]{}
	args := make(map[string]interface{})
	args["id_"] = id
	call := WheelCall[KeyGen]{"key.gen", args, result}
	return call
}

func KeyGenAcceptSync(id string, force bool) WheelCall[KeyGen] {
	var result = &APIReturnType[WheelResultType[KeyGen]]{}
	args := make(map[string]interface{})
	args["id_"] = id
	args["force"] = force
	call := WheelCall[KeyGen]{"key.gen_accept", args, result}
	return call
}
