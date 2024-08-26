/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/10 11:22:46 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/map/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/
// qmemcodestart
package map000

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/gookit/color"

	commonprint "github.com/gainovel/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

func ConcurrentMapReads() {
	var (
		m1                   map[int]int
		wg                   sync.WaitGroup
		n                    int
		targetKey, targetVal int
	)
	n = 1000
	m1 = make(map[int]int, n)

	// 先随机添加一些元素值
	for i := 0; i < n; i++ {
		key := rand.Intn(i+1000) + 100
		val := rand.Intn(key) + 88
		m1[key] = val
	}

	targetKey = 100
	targetVal = 10000
	m1[targetKey] = targetVal
	// 开启1000个协程并发读
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) {
			defer wg.Done()
			color.HiMagenta.Printf("goroutine %d read start...\n", j)
			time.Sleep(time.Microsecond * 10)
			val := m1[targetKey]
			fmt.Printf("groutine %s get targetKey(%d), targetVal(%d)\n", color.HiCyan.Sprintf("%d", j), targetKey, val)
		}(i)
	}
	wg.Wait()
}
func ConcurrentMapWrites() {
	var (
		m1                   map[int]int
		wg                   sync.WaitGroup
		n                    int
		targetKey, targetVal int
	)
	n = 1000
	m1 = make(map[int]int, n)

	// 先随机添加一些元素值
	for i := 0; i < n; i++ {
		key := rand.Intn(i+1000) + 100
		val := rand.Intn(key) + 88
		m1[key] = val
	}

	targetKey = 1000
	targetVal = 10000

	// 开启1000个协程并发写
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) {
			defer wg.Done()
			color.HiMagenta.Printf("goroutine %d write start...\n", j)
			//fmt.Printf("groutine %s write targetKey(%d), targetVal(%d)\n", color.New(color.FgHiCyan).Sprintf("%d", j), targetKey, targetVal)
			time.Sleep(time.Microsecond * 10)
			m1[targetKey] = targetVal
		}(i)
	}
	wg.Wait()
}

func TestName_2024_01_10_11_22_46(t *testing.T) {
	// 多协程同时写会报错concurrent map writes
	// go test -v -run TestName_2024_01_10_11_22_46/concurrent_map_write github.com/gainovel/testcase/stdlib/runtime/map
	t.Run("concurrent map write", func(t *testing.T) {
		// 使用windows terminal 在./cmd/main.go测试ConcurrentMapWrites()
		ConcurrentMapWrites()
	})
	// panic的情况👉给nil map添加key-value
	// go test -v -run TestName_2024_01_10_11_22_46/assignment_to_entry_in_nil_map github.com/gainovel/testcase/stdlib/runtime/map
	t.Run("assignment to entry in nil map", func(t *testing.T) {
		var (
			m1 map[int]int
		)
		m1[1] = 2
	})

	// map 的简单使用
	// 使用内置函数delete()进行删除
	// 查询map时，使用逗号模式(val,ok)获取值，避免操作零值，ok表示key是否存在
	// go test -v -run TestName_2024_01_10_11_22_46/map_crud github.com/gainovel/testcase/stdlib/runtime/map
	t.Run("map crud", func(t *testing.T) {
		var (
			m1  map[int]int
			key int
			ok  bool
		)
		m1 = make(map[int]int)
		m1[1] = 101
		m1[2] = 102
		myfmt.VarInitPrintln(`var (
	m1 map[int]int
)`)
		myfmt.ColorDescPrintln("m1 = make(map[int]int)", "m1[1] = 101", "m1[2] = 102")
		myfmt.KeyValuePrintln("m1", m1, "len(m1)", len(m1))
		delete(m1, 1)
		myfmt.ColorDescPrintln("delete(m1, 1)")
		myfmt.KeyValuePrintln("m1", m1, "len(m1)", len(m1))
		key, ok = m1[1]
		myfmt.ColorDescPrintln("key,ok = m1[1]")
		myfmt.KeyValuePrintln("m1", m1, "len(m1)", len(m1), "key", key, "ok", ok)
		key, ok = m1[2]
		myfmt.ColorDescPrintln("key,ok = m1[2]")
		myfmt.KeyValuePrintln("m1", m1, "len(m1)", len(m1), "key", key, "ok", ok)
	})
	t.Run("nil map", func(t *testing.T) {
		var (
			m1 map[int]int
		)
		m1[1] = 2
		fmt.Println(m1[1])
	})
	t.Run("map address", func(t *testing.T) {
		var (
			m1 = make(map[int]int)
		)
		m1[1] = 2
	})
}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd docs/tests/stdlib/runtime/map.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/runtime/map/features_usages_001_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/concurrent_map_write -f Makefiles/stdlib/runtime/map.mk
//*  Test Result: fatal error: concurrent map writes
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/assignment_to_entry_in_nil_map -f Makefiles/stdlib/runtime/map.mk
//*  Test Result: panic: assignment to entry in nil map
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/map_crud -f Makefiles/stdlib/runtime/map.mk
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*  var (
//*          m1 map[int]int
//*  )
//*  --------------------------------------------------------------------------------
//*  👇
//*  m1 = make(map[int]int)
//*  m1[1] = 101
//*  m1[2] = 102
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                m1 | map[1:101 2:102]
//*  |           len(m1) | 2
//*  --------------------------------------------------------------------------------
//*  👇
//*  delete(m1, 1)
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                m1 | map[2:102]
//*  |           len(m1) | 1
//*  --------------------------------------------------------------------------------
//*  👇
//*  key,ok = m1[1]
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                m1 | map[2:102]
//*  |           len(m1) | 1
//*  |               key | 0
//*  |                ok | false
//*  --------------------------------------------------------------------------------
//*  👇
//*  key,ok = m1[2]
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                m1 | map[2:102]
//*  |           len(m1) | 1
//*  |               key | 102
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
// qmemoutputend
