package main

import (
	"fmt"
	"github.com/rjmateus/go-salt-api-client/calls/wheel"
	"github.com/rjmateus/go-salt-api-client/client"
)

func main4() {

	/*clientSalt := client.NewClient("https://m43-server.tf.local:9080",
	"admin", "328ae1ec5a83f4b6df81dc8392b0b8fa1367c910efd22a6d978c735b97061e3f",
	"file")*/

	clientSalt := client.NewClient("https://m43-server.tf.local:9080",
		"admin", "328ae1ec5a83f4b6df81dc8392b0b8fa1367c910efd22a6d978c735b97061e3f",
		"file")

	args := make([]interface{}, 0)
	args = append(args, "ls")

	keys := wheel.KeyListAll().CallSync(clientSalt)

	fmt.Printf("KEYS data: >%s<\n", keys)
	for _, val := range keys.Data.Return.MinionsPre {
		fmt.Printf("accepting %s\n", val)
		//wheel.KeyAccept(clientSalt, val)
	}

	wheel.KeyGen("my_id1").CallSync(clientSalt)
	wheel.KeyGenAccept("my-id-2", true).CallSync(clientSalt)

	keyMinion := wheel.KeyFinger("m43-minion-suse.tf.local").CallSync(clientSalt)
	fmt.Printf("Finger Minion: >%s<\n", keyMinion)
}
