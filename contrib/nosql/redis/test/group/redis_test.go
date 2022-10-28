// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package group_test

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx      = gctx.GetInitCtx()
	redis, _ = gredis.New(&gredis.Config{
		Address: `:6379`,
		Db:      1,
	})
)
