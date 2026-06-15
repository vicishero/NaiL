//go:build constraint

package security

import (
	"github.com/vicishero/NaiL/server/internal/core"
)

var (
	_ core.AttachmentCheckService = (*attachmentCheckServant)(nil)
	_ core.PhoneVerifyService     = (*juheSmsServant)(nil)
)
