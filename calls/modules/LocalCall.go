package modules

import (
	"github.com/rjmateus/go-salt-api-client/client"
	"log"
)

func (w LocalCall) CallSync(saltClient client.SaltClient) []byte {
	data, err := saltClient.Call(w, client.LOCAL, w.Target)
	if err != nil {
		log.Printf("error calling %s-> %s", w.Fun, err)
	}
	return data
}

func (w LocalCall) GetPayload() map[string]interface{} {
	payload := make(map[string]interface{})
	payload["fun"] = w.Fun
	if w.Args != nil {
		payload["arg"] = w.Args
	}
	if w.Kwargs != nil {
		for k, v := range w.Kwargs {
			payload[k] = v
		}
	}

	return payload
}
