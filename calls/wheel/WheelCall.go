package wheel

import (
	"encoding/json"
	"github.com/rjmateus/go-salt-api-client/client"
	"log"
)

type WheelCall[T any] struct {
	Fun    string
	Kwargs map[string]interface{}
	result *APIReturn[WheelResult[T]]
}

func (w WheelCall[T]) CallSync(saltClient client.SaltClient) WheelResult[T] {
	data, err := saltClient.Call(w, client.WHEEL, nil)
	if err != nil {
		log.Printf("error calling %s-> %s", w.Fun, err)
	}
	json.Unmarshal(data, w.result)
	// FIXME check data is present before returning it
	return w.result.Return[0]
}

func (w WheelCall[T]) GetPayload() map[string]interface{} {
	payload := make(map[string]interface{})
	payload["fun"] = w.Fun
	if w.Kwargs != nil {
		for k, v := range w.Kwargs {
			payload[k] = v
		}
	}
	return payload
}
