//go:build constraint

package space

import (
	api "github.com/vicishero/NaiL/server/auto/api/x/v1"
)

var _ api.User = (*userSrv)(nil)
