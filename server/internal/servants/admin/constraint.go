//go:build constraint

package admin

import (
	api "github.com/vicishero/NaiL/server/auto/api/m/v1"
)

var _ api.User = (*userSrv)(nil)
