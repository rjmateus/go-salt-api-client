package calls

type Call interface {
	GetPayload() map[string]interface{}
}
