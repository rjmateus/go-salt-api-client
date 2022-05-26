package wheel

type KeysAll struct {
	Minions         []string `json:"minions"`
	MinionsPre      []string `json:"minions_pre"`
	MinionsRejected []string `json:"minions_rejected"`
	MinionsDenied   []string `json:"minions_denied"`
	Local           []string `json:"local"`
}

type WheelDataType[T any] struct {
	Fun    string `json:"fun"`
	Jid    string `json:"jid"`
	User   string `json:"user"`
	Stamp  string `json:"_stamp"`
	Return T      `json:"return"`
}

type WheelResultType[T any] struct {
	Tag  string           `json:"tag"`
	Data WheelDataType[T] `json:"data"`
}

type APIReturnType[T any] struct {
	Return []WheelResultType[T] `json:"return"`
}
