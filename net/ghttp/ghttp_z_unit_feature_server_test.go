// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp_test

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
	"time"
)

type ghttp_test_plugin struct {
}

func (g *ghttp_test_plugin) Name() string {
	return "ghttp_test_plugin Name"
}

func (g *ghttp_test_plugin) Author() string {
	return "ghttp_test_plugin Author"
}

func (g *ghttp_test_plugin) Version() string {
	return "ghttp_test_plugin Version"
}

func (g *ghttp_test_plugin) Description() string {
	return "ghttp_test_plugin Description"
}

func (g *ghttp_test_plugin) Install(s *ghttp.Server) error {
	return nil
}

func (g *ghttp_test_plugin) Remove() error {
	return nil
}

func Test_Server_GetServer(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := ghttp.GetServer("GoFrame")
		t.AssertNE(s, nil)
		gf := ghttp.GetServer("GoFrame")
		t.Assert(s, gf)

		t.AssertNE(s.GetOpenApi(), nil)
	})
}

func Test_Server_Run(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := ghttp.GetServer("GoFrame")
		t.AssertNE(s, nil)
		s.Plugin(&ghttp_test_plugin{})
		go s.Run()
		time.Sleep(time.Millisecond * 500)
		t.Assert(s.Status(), ghttp.ServerStatusRunning)
	})
}

func Test_Server_Wait(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := ghttp.GetServer("1")
		t.AssertNE(s1, nil)
		s1.Plugin(&ghttp_test_plugin{})
		s1.Start()
		defer s1.Shutdown()

		s2 := ghttp.GetServer("2")
		t.AssertNE(s2, nil)
		s1.Plugin(&ghttp_test_plugin{})
		s2.Start()
		defer s2.Shutdown()

		go func() {
			time.Sleep(time.Millisecond * 2000)
			ghttp.ShutdownAllServer(ctx)
		}()

		ghttp.Wait()
	})
}
