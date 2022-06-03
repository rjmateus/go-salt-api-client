package main

import (
	"fmt"
	"github.com/rjmateus/go-salt-api-client/calls/runner"
	"github.com/rjmateus/go-salt-api-client/client"
)

func main() {

	clientSalt := client.NewClient("https://m42-server.tf.local:9080",
		"admin", "78b4c3f73ab2750e7c84a4aa028002c24d72e9119f46696e8eb9b0f4156fd376",
		"file")

	args := make([]interface{}, 0)
	args = append(args, "ls")

	call := runner.LookupJid("20220603145550816135")

	result := call.CallSync(clientSalt)
	fmt.Println(result)

}
