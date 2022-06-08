package wheel

/*
* More information
* https://docs.saltproject.io/en/latest/ref/wheel/all/salt.wheel.key.html
 */

func KeyListAll() WheelCall[Keys[[]string]] {
	var result = &APIReturnType[WheelResultType[Keys[[]string]]]{}
	call := WheelCall[Keys[[]string]]{"key.list_all", nil, result}
	return call
}

func KeyFinger(match string) WheelCall[Keys[map[string]string]] {
	var result = &APIReturnType[WheelResultType[Keys[map[string]string]]]{}

	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[Keys[map[string]string]]{"key.finger", args, result}
	return call
}

func KeyAccept(match string) WheelCall[interface{}] {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.accept", args, result}
	return call
}

func KeyReject(match string) WheelCall[interface{}] {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.reject", args, result}
	return call
}

func KeyDelete(match string) WheelCall[interface{}] {
	var result = &APIReturnType[WheelResultType[interface{}]]{}
	args := make(map[string]interface{})
	args["match"] = match
	call := WheelCall[interface{}]{"key.delete", args, result}
	return call
}

func KeyGen(id string) WheelCall[KeyGenResult] {
	var result = &APIReturnType[WheelResultType[KeyGenResult]]{}
	args := make(map[string]interface{})
	args["id_"] = id
	call := WheelCall[KeyGenResult]{"key.gen", args, result}
	return call
}

func KeyGenAccept(id string, force bool) WheelCall[KeyGenResult] {
	var result = &APIReturnType[WheelResultType[KeyGenResult]]{}
	args := make(map[string]interface{})
	args["id_"] = id
	args["force"] = force
	call := WheelCall[KeyGenResult]{"key.gen_accept", args, result}
	return call
}
