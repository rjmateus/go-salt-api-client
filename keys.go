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

	clientSalt := client.NewClient("https://m42-server.tf.local:9080",
		"admin", "78b4c3f73ab2750e7c84a4aa028002c24d72e9119f46696e8eb9b0f4156fd376",
		"file")

	args := make([]interface{}, 0)
	args = append(args, "ls")

	keys := wheel.KeyListAllSync(clientSalt)

	fmt.Printf("KEYS data: >%s<\n", keys)
	for _, val := range keys.MinionsPre {
		fmt.Println("accepting %s", val)
		//wheel.KeyAcceptSync(clientSalt, val)
	}

	wheel.KeyGenSync(clientSalt, "my_id1")
	wheel.KeyGenAcceptSync(clientSalt, "my-id-2", true)

	keyMinion := wheel.KeyFingerSync(clientSalt, "m43-minion-suse.tf.local")
	fmt.Printf("Finger Minion: >%s<\n", keyMinion)

}
