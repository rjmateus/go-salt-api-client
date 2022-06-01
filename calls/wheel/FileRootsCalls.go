package wheel

import "github.com/rjmateus/go-salt-api-client/client"

func FileRootsListRoot(saltClient client.SaltClient) WheelResultType[FileRootListRoot] {
	var result = &APIReturnType[WheelResultType[FileRootListRoot]]{}
	call := WheelCall[FileRootListRoot]{"file_roots.list_roots", nil, result}
	return call.CallSync(saltClient)
}

func FileRootsListEnv(saltClient client.SaltClient, saltenv string) WheelResultType[FileRootListRoot] {
	var result = &APIReturnType[WheelResultType[FileRootListRoot]]{}
	args := make(map[string]interface{})
	if len(saltenv) > 0 {
		args[saltenv] = saltenv
	}
	call := WheelCall[FileRootListRoot]{"file_roots.list_env", args, result}
	return call.CallSync(saltClient)
}
