package runner

type RunnerCall[T any] struct {
	Fun    string
	Kwargs map[string]interface{}
	result *T
}
