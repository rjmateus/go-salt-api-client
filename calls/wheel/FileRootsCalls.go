package wheel

/**
* https://docs.saltproject.io/en/latest/ref/wheel/all/salt.wheel.file_roots.html
 */

func FileRootsListRoot() WheelCall[FileRootListRoot] {
	var result = &APIReturn[WheelResult[FileRootListRoot]]{}
	call := WheelCall[FileRootListRoot]{"file_roots.list_roots", nil, result}
	return call
}

func FileRootsListEnv(saltenv string) WheelCall[FileRootListRoot] {
	var result = &APIReturn[WheelResult[FileRootListRoot]]{}
	args := make(map[string]interface{})
	if len(saltenv) > 0 {
		args[saltenv] = saltenv
	}
	call := WheelCall[FileRootListRoot]{"file_roots.list_env", args, result}
	return call
}
