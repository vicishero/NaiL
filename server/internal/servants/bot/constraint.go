//go:build constraint

package bot

import (
	api "github.com/vicishero/NaiL/server/auto/api/r/v1"
)

var _ api.User = (*userSrv)(nil)
