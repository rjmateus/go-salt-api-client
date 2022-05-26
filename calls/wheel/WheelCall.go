package wheel

import (
	"github.com/rjmateus/go-salt-api-client/client"
	"log"
)

type WheelCall struct {
	Fun    string
	Kwargs map[string]interface{}
	// TODO have a look in generics instead of having a parser
	// https://go.dev/doc/tutorial/generics
	Parser func([]byte) interface{}
}

func (w WheelCall) CallSync(saltClient client.SaltClient) interface{} {
	result, err := saltClient.Call(w, client.WHEEL, nil)
	if err != nil {
		log.Printf("error calling %s-> %s", w.Fun, err)
	}
	return result
}
func (w WheelCall) GetParser() func([]byte) interface{} {
	return w.Parser
}
func (w WheelCall) GetPayload() map[string]interface{} {
	payload := make(map[string]interface{})
	payload["fun"] = w.Fun
	if w.Kwargs != nil {
		for k, v := range w.Kwargs {
			payload[k] = v
		}
	}
	return payload
}
