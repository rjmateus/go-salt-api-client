package client

type ClientType string

const (
	LOCAL        ClientType = "local"
	LOCAL_ASYNC             = "local_async"
	LOCAL_BATCH             = "local_batch"
	RUNNER                  = "runner"
	RUNNER_ASYNC            = "runner_async"
	SSH                     = "ssh"
	WHEEL                   = "wheel"
	WHEEL_ASYNC             = "wheel_async"
)
