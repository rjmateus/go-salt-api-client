package main

import (
	"fmt"
	"github.com/rjmateus/go-salt-api-client/calls/modules/calls"
	"github.com/rjmateus/go-salt-api-client/calls/wheel"
	"github.com/rjmateus/go-salt-api-client/client"
)

func main() {

	/*clientSalt := client.NewClient("https://m43-server.tf.local:9080",
	"admin", "328ae1ec5a83f4b6df81dc8392b0b8fa1367c910efd22a6d978c735b97061e3f",
	"file")*/

	clientSalt := client.NewClient("https://m42-server.tf.local:9080",
		"admin", "78b4c3f73ab2750e7c84a4aa028002c24d72e9119f46696e8eb9b0f4156fd376",
		"file")

	args := make([]interface{}, 0)
	args = append(args, "ls")

	keys := wheel.KeyListAllSync().CallSync(clientSalt)

	fmt.Printf("KEYS data: >%s<\n", keys)
	for _, val := range keys.Data.Return.MinionsPre {
		fmt.Printf("accepting %s\n", val)
		//wheel.KeyAcceptSync(clientSalt, val)
	}

	wheel.KeyGenSync("my_id1").CallSync(clientSalt)
	wheel.KeyGenAcceptSync("my-id-2", true).CallSync(clientSalt)

	keyMinion := wheel.KeyFingerSync("m43-minion-suse.tf.local").CallSync(clientSalt)
	fmt.Printf("Finger Minion: >%s<\n", keyMinion)

	data := calls.ListMinions(clientSalt)
	fmt.Printf("MINIONS: >%s<\n", data)

	tt := wheel.FileRootsListEnv("").CallSync(clientSalt)
	fmt.Println(tt)

}
