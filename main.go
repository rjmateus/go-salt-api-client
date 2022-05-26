package main

import (
	"fmt"
	"github.com/rjmateus/go-salt-api-client/calls/wheel"
	"github.com/rjmateus/go-salt-api-client/client"
	"github.com/rjmateus/go-salt-api-client/target"
	"log"
)

func main() {

	clientSalt := client.NewClient("https://m43-server.tf.local:9080",
		"admin", "328ae1ec5a83f4b6df81dc8392b0b8fa1367c910efd22a6d978c735b97061e3f",
		"file")

	args := make([]interface{}, 0)
	args = append(args, "ls")

	data, err := clientSalt.Run(client.LOCAL, "cmd.run", target.NewTargetGlob("*"), args, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("decode data: >%s<\n", data)

	data, err = clientSalt.Run(client.LOCAL, "test.ping", target.NewTargetList([]string{"m43-minion-suse.tf.local"}), nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("decode data: >%s<\n", data)

	data, err = clientSalt.Run(client.LOCAL_ASYNC, "test.ping", target.NewTargetGrains("os", "SUSE"), nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("decode data: >%s<\n", data)

	keys := wheel.KeyListAllSync(clientSalt)

	fmt.Printf("KEYS data: >%s<\n", keys)

	keyMinion := wheel.KeyFinger("m43-minion-suse.tf.local").CallSync(clientSalt)
	fmt.Printf("KEY Minion: >%s<\n", keyMinion)
}
