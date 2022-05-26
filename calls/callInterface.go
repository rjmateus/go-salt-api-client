package calls

type Call interface {
	GetPayload() map[string]interface{}
	GetParser() func([]byte) interface{}
}
