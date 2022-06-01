package calls

import (
	"github.com/rjmateus/go-salt-api-client/calls/modules"
	"github.com/rjmateus/go-salt-api-client/client"
)

func ListMinions(saltClient client.SaltClient) []byte {
	call := modules.LocalCall{"minion.list", nil, nil, nil}
	return call.CallSync(saltClient)
}
