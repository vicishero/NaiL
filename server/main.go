// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/vicishero/NaiL/server/cmd"
	_ "github.com/vicishero/NaiL/server/cmd/migrate"
	_ "github.com/vicishero/NaiL/server/cmd/serve"
)

func main() {
	cmd.Execute()
}
