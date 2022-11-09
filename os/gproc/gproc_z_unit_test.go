// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*" -benchmem

package gproc_test

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_ShellExec(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s, err := gproc.ShellExec(gctx.New(), `echo 123`)
		t.AssertNil(err)
		t.Assert(s, "123\n")
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		_, err := gproc.ShellExec(gctx.New(), `NoneExistCommandCall`)
		t.AssertNE(err, nil)
	})
}

func TestPid(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertGE(gproc.Pid(), 0)
	})
}

func TestPPid(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertGE(gproc.PPid(), 0)
		t.Assert(gproc.SetPPid(0), nil)
		t.AssertGT(gproc.PPid(), 0)
		t.Assert(gproc.SetPPid(12345), nil)
		t.Assert(gproc.PPid(), 12345)
	})
}

func TestPPidOS(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertGE(gproc.PPidOS(), 0)
	})
}

func TestIsChild(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gproc.IsChild(), true)
	})
}

func TestStartTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertGE(gproc.StartTime(), 0)
		t.AssertGE(gproc.Uptime(), 0)
	})
}

func TestSearchBinary(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		filePath := "./gproc.go"
		t.Assert(gproc.SearchBinary("./gproc.go"), filePath)
		t.Assert(gproc.SearchBinary("./noexist.go"), "")
	})
}
