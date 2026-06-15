//go:build constraint

package mobile

import (
	api "github.com/vicishero/NaiL/server/auto/rpc/greet/v1"
)

var _ api.GreetServiceServer = (*greetServiceSrv)(nil)
