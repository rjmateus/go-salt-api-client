package modules

import "github.com/rjmateus/go-salt-api-client/target"

type LocalCall struct {
	Fun    string
	Args   []interface{}
	Kwargs map[string]interface{}
	Target *target.SaltTarget
}
