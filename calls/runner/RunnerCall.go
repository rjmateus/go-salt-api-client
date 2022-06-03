package runner

import (
	"encoding/json"
	"github.com/rjmateus/go-salt-api-client/client"
	"log"
)

func (w RunnerCall[T]) CallSync(saltClient client.SaltClient) *T {
	data, err := saltClient.Call(w, client.RUNNER, nil)
	if err != nil {
		log.Printf("error calling %s-> %s", w.Fun, err)
	}
	json.Unmarshal(data, w.result)
	// FIXME check data is present before returning it
	return w.result
}

func (w RunnerCall[T]) GetPayload() map[string]interface{} {
	payload := make(map[string]interface{})
	payload["fun"] = w.Fun
	if w.Kwargs != nil {
		for k, v := range w.Kwargs {
			payload[k] = v
		}
	}
	return payload
}
