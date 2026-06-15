//go:build constraint

package localoss

import (
	api "github.com/vicishero/NaiL/server/auto/api/s/v1"
)

var _ api.User = (*userSrv)(nil)
