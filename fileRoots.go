package main

import (
	"fmt"
	"github.com/rjmateus/go-salt-api-client/calls/wheel"
	"github.com/rjmateus/go-salt-api-client/client"
)

func main() {

	/*clientSalt := client.NewClient("https://m43-server.tf.local:9080",
	"admin", "328ae1ec5a83f4b6df81dc8392b0b8fa1367c910efd22a6d978c735b97061e3f",
	"file")*/

	clientSalt := client.NewClient("https://m43-server.tf.local:9080",
		"admin", "328ae1ec5a83f4b6df81dc8392b0b8fa1367c910efd22a6d978c735b97061e3f",
		"file")

	args := make([]interface{}, 0)
	args = append(args, "ls")

	//tt := wheel.FileRootsListRoot().CallSync(clientSalt)
	//fmt.Println(tt)

	ttAsync := wheel.FileRootsListRoot().CallAsync(clientSalt)
	fmt.Println(ttAsync)
}
