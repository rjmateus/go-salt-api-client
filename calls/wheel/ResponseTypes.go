package wheel

/**
* Generic wheel response type
 */

type APIReturn[R any] struct {
	Return []R `json:"return"`
}

type WheelResult[T any] struct {
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

/**
* Key datatype
 */

type Keys[T any] struct {
	Minions         T `json:"minions"`
	MinionsPre      T `json:"minions_pre"`
	MinionsRejected T `json:"minions_rejected"`
	MinionsDenied   T `json:"minions_denied"`
	Local           T `json:"local"`
}

type KeyGenResult struct {
	Pub  string
	Priv string
}

/**
* File Root datatype
 */

type FileRootListRoot struct {
	Base []map[string]interface{}
}

type WheelAsyncResult struct {
	Tag string `json:"tag"`
	Jid string `json:"jid"`
}
