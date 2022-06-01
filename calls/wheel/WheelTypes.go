package wheel

type WheelCall[T any] struct {
	Fun    string
	Kwargs map[string]interface{}
	result *APIReturnType[WheelResultType[T]]
}

type APIReturnType[R any] struct {
	Return []R `json:"return"`
}

type WheelResultType[T any] struct {
	Tag  string           `json:"tag"`
	Data WheelDataType[T] `json:"data"`
}

type WheelDataType[T any] struct {
	Fun    string `json:"fun"`
	Jid    string `json:"jid"`
	User   string `json:"user"`
	Stamp  string `json:"_stamp"`
	Return T      `json:"return"`
}

type Keys[T any] struct {
	Minions         T `json:"minions"`
	MinionsPre      T `json:"minions_pre"`
	MinionsRejected T `json:"minions_rejected"`
	MinionsDenied   T `json:"minions_denied"`
	Local           T `json:"local"`
}

type KeyGen struct {
	Pub  string
	Priv string
}

type FileRootListRoot struct {
	Base []map[string]interface{}
}
