package main

import (
	"fmt"
	"github.com/rjmateus/go-salt-api-client/calls/runner"
	"github.com/rjmateus/go-salt-api-client/client"
)

func main3() {

	clientSalt := client.NewClient("https://m43-server.tf.local:9080",
		"admin", "328ae1ec5a83f4b6df81dc8392b0b8fa1367c910efd22a6d978c735b97061e3f",
		"file")

	args := make([]interface{}, 0)
	args = append(args, "ls")

	call := runner.LookupJid("20220603145550816135")

	result := call.CallSync(clientSalt)
	fmt.Println(result)

}
