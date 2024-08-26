/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/11 16:18:49 星期四
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/sync/map/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package syncmap000

import (
	"fmt"
	commontools "github.com/gainovel/testcase/tools/common"
	"sync"
	"testing"
)

func TestName_2024_01_11_16_18_49(t *testing.T) {
	//  go test -v -run TestName_2024_01_11_16_18_49/sync.Map_1 github.com/gainovel/testcase/stdlib/sync/map
	t.Run("sync.Map 1", func(t *testing.T) {
		var (
			m1   sync.Map
			val  any
			ok   bool
			temp string
		)
		temp = "{}"
		fmt.Println()
		commontools.PrintAll(true, "var m1  sync.Map", "sync.Map m1 status", "", "sync.Map m1", temp)
		m1.Store("Jim", 80)
		m1.Store("Kevin", 90)
		m1.Store("Jane", 100)
		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "add "+temp+"to sync.Map m1", "sync.Map m1 status", "", "sync.Map m1", temp)
		val, ok = m1.Load("Jim")
		commontools.PrintAll(true, "m1.Load(\"Jim\")", "Load Result", "", "val", val, "ok", ok)
		m1.Delete("Jim")
		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "m1.Delete(\"Jim\")", "sync.Map m1 status", "", "sync.Map m1", temp)
	})
	// go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadOrStore github.com/gainovel/testcase/stdlib/sync/map
	t.Run("sync.Map.LoadOrStore", func(t *testing.T) {
		var (
			m1     sync.Map
			actual any
			loaded bool
			temp   string
		)
		temp = "{}"
		fmt.Println()
		commontools.PrintAll(true, "var m1  sync.Map", "sync.Map m1 status", "", "sync.Map m1", temp)
		m1.Store("Jim", 80)
		m1.Store("Kevin", 90)
		m1.Store("Jane", 100)
		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "add "+temp+"to sync.Map m1", "sync.Map m1 status", "", "sync.Map m1", temp)
		actual, loaded = m1.LoadOrStore("Jim", 81)
		commontools.PrintAll(true, "m1.LoadOrStore(\"Jim\", 81)", "LoadOrStore Result", "", "actual", actual, "loaded", loaded)
		m1.Delete("Jim")
		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "m1.Delete(\"Jim\")", "sync.Map m1 status", "", "sync.Map m1", temp)
		actual, loaded = m1.LoadOrStore("Jim", 81)
		commontools.PrintAll(true, "m1.LoadOrStore(\"Jim\", 81)", "LoadOrStore Result", "", "actual", actual, "loaded", loaded)
	})
	// go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadAndDelete github.com/gainovel/testcase/stdlib/sync/map
	t.Run("sync.Map.LoadAndDelete", func(t *testing.T) {
		var (
			m1     sync.Map
			value  any
			loaded bool
			temp   string
		)
		temp = "{}"
		fmt.Println()
		commontools.PrintAll(true, "var m1  sync.Map", "sync.Map m1 status", "", "sync.Map m1", temp)
		m1.Store("Jim", 80)
		m1.Store("Kevin", 90)
		m1.Store("Jane", 100)

		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "add "+temp+"to sync.Map m1", "sync.Map m1 status", "", "sync.Map m1", temp)

		value, loaded = m1.LoadAndDelete("Jim")
		commontools.PrintAll(true, "m1.LoadAndDelete(\"Jim\")", "LoadAndDelete Result", "", "value", value, "loaded", loaded)

		temp = ""
		m1.Range(func(key, value any) bool {
			temp += fmt.Sprintf("{%v:%v} ", key, value)
			return true
		})
		commontools.PrintAll(true, "m1.Range", "sync.Map m1 status", "", "sync.Map m1", temp)

		value, loaded = m1.LoadAndDelete("Jim")
		commontools.PrintAll(true, "m1.LoadAndDelete(\"Jim\")", "LoadAndDelete Result", "", "value", value, "loaded", loaded)
	})
}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd docs/tests/stdlib/sync/map.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/sync/map/features_usages_001_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/sync.Map_1 -f Makefiles/stdlib/sync/map.mk
//*  Test Result:
//*   (☀️2024年03月11日🕛下午 2点37分18秒)👇(var m1  sync.Map)
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {}
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点37分18秒)👇(add {Jim:80} {Kevin:90} {Jane:100} to sync.Map m1)
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {Jim:80} {Kevin:90} {Jane:100}
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点37分18秒)👇(m1.Load("Jim"))
//*  ----------------------------------------------------------------------------------------
//*  | Load Result         |
//*  ----------------------------------------------------------------------------------------
//*  | val                 |  80
//*  | ok                  |  true
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点37分18秒)👇(m1.Delete("Jim"))
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {Kevin:90} {Jane:100}
//*  ----------------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/sync.Map.LoadOrStore -f Makefiles/stdlib/sync/map.mk
//*  Test Result:
//*   (☀️2024年03月11日🕛下午 2点44分40秒)👇(var m1  sync.Map)
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {}
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点44分40秒)👇(add {Jane:100} {Jim:80} {Kevin:90} to sync.Map m1)
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {Jane:100} {Jim:80} {Kevin:90}
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点44分40秒)👇(m1.LoadOrStore("Jim", 81))
//*  ----------------------------------------------------------------------------------------
//*  | LoadOrStore Result  |
//*  ----------------------------------------------------------------------------------------
//*  | actual              |  80
//*  | loaded              |  true
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点44分40秒)👇(m1.Delete("Jim"))
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {Kevin:90} {Jane:100}
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点44分40秒)👇(m1.LoadOrStore("Jim", 81))
//*  ----------------------------------------------------------------------------------------
//*  | LoadOrStore Result  |
//*  ----------------------------------------------------------------------------------------
//*  | actual              |  81
//*  | loaded              |  false
//*  ----------------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/sync.Map.LoadAndDelete -f Makefiles/stdlib/sync/map.mk
//*  Test Result:
//*   (☀️2024年03月11日🕛下午 2点45分08秒)👇(var m1  sync.Map)
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {}
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点45分08秒)👇(add {Jim:80} {Kevin:90} {Jane:100} to sync.Map m1)
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {Jim:80} {Kevin:90} {Jane:100}
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点45分08秒)👇(m1.LoadAndDelete("Jim"))
//*  ----------------------------------------------------------------------------------------
//*  | LoadAndDelete Result|
//*  ----------------------------------------------------------------------------------------
//*  | value               |  80
//*  | loaded              |  true
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点45分08秒)👇(m1.Range)
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1 status  |
//*  ----------------------------------------------------------------------------------------
//*  | sync.Map m1         |  {Jane:100} {Kevin:90}
//*  ----------------------------------------------------------------------------------------
//*     (☀️2024年03月11日🕛下午 2点45分08秒)👇(m1.LoadAndDelete("Jim"))
//*  ----------------------------------------------------------------------------------------
//*  | LoadAndDelete Result|
//*  ----------------------------------------------------------------------------------------
//*  | value               |  <nil>
//*  | loaded              |  false
//*  ----------------------------------------------------------------------------------------
//**************************************************************************************
//
// qmemoutputend
