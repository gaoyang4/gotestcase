/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/10 16:36:44 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/iota/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package iota000

import (
	"fmt"
	"testing"

	commonprint "github.com/gainovel/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

func TestName_2024_01_10_16_36_44(t *testing.T) {
	// ‼️ iota 和 const声明块 关键两条原则，所有赋值不疑惑
	// 1.iota代表了const声明块的行索引
	// 2.如果为常量指定了一个表达式，但后续常量没有表达式，则继承上面的表达式
	// go test -v -run TestName_2024_01_10_16_36_44/iota_1 github.com/gainovel/testcase/stdlib/runtime/iota
	t.Run("iota 1", func(t *testing.T) {
		const (
			a, b = 1 << iota, iota * 2 //【iota:0】 【1<<iota:1】 【iota*2:0】
			c, d                       //【iota:1】 【1<<iota:2】 【iota*2:2】
			e, f                       //【iota:2】 【1<<iota:4】 【iota*2:4】
			g, h                       //【iota:3】 【1<<iota:8】 【iota*2:6】
			i, j                       //【iota:4】 【1<<iota:16】 【iota*2:8】
			k, l                       //【iota:5】 【1<<iota:32】 【iota*2:10】
		)
		myfmt.VarInitPrintln(`const (
	a, b = 1 << iota, iota * 2 //【iota:0】 【1<<iota:1】 【iota*2:0】
	c, d                       //【iota:1】 【1<<iota:2】 【iota*2:2】
	e, f                       //【iota:2】 【1<<iota:4】 【iota*2:4】
	g, h                       //【iota:3】 【1<<iota:8】 【iota*2:6】
	i, j                       //【iota:4】 【1<<iota:16】 【iota*2:8】
	k, l                       //【iota:5】 【1<<iota:32】 【iota*2:10】
)`)
		myfmt.KeyValuePrintln("a", a, "b", b, "c", c, "d", d, "e", e, "f", f, "g", g, "h", h, "i", i, "j", j, "k", k, "l", l)

	})
	// go test -v -run TestName_2024_01_10_16_36_44/iota_2 github.com/gainovel/testcase/stdlib/runtime/iota
	t.Run("iota 2", func(t *testing.T) {
		// const声明块即使常量中途重新声明表达式，iota的值始终都是该行的行索引
		const (
			a, b = 1 << iota, iota * 2 //【iota:0】 【1<<iota:1】 【iota*2:0】
			c, d                       //【iota:1】 【1<<iota:2】 【iota*2:2】
			e, f = iota, iota + 2      //【iota:2】 【iota:2】 【iota+2:4】
			g, h                       //【iota:3】 【iota:3】 【iota+2:5】
			i, j = iota * 3, iota + 4  //【iota:4】 【iota*3:12】 【iota+4:8】
			k, l                       //【iota:5】 【iota*3:15】 【iota+4:9】
		)
		myfmt.VarInitPrintln(`const (
	a, b = 1 << iota, iota * 2 //【iota:0】 【1<<iota:1】 【iota*2:0】
	c, d                       //【iota:1】 【1<<iota:2】 【iota*2:2】
	e, f = iota, iota + 2      //【iota:2】 【iota:2】 【iota+2:4】
	g, h                       //【iota:3】 【iota:3】 【iota+2:5】
	i, j = iota * 3, iota + 4  //【iota:4】 【iota*3:12】 【iota+4:8】
	k, l                       //【iota:5】 【iota*3:15】 【iota+4:9】
)`)
		myfmt.KeyValuePrintln("a", a, "b", b, "c", c, "d", d, "e", e, "f", f, "g", g, "h", h, "i", i, "j", j, "k", k, "l", l)

	})
	// go test -v -run TestName_2024_01_10_16_36_44/iota_3 github.com/gainovel/testcase/stdlib/runtime/iota
	t.Run("iota 3", func(t *testing.T) {
		// const声明块即使首行非iota表达式，iota的值始终都是该行的行索引
		const (
			x, y = 100, 100
			a, b = 1 << iota, iota * 2 //【iota:1】 【1<<iota:2】 【iota*2:2】
			c, d                       //【iota:2】 【1<<iota:4】 【iota*2:4】
			e, f = iota, iota + 2      //【iota:3】 【iota:3】 【iota+2:5】
			g, h                       //【iota:4】 【iota:4】 【iota+2:6】
		)
		myfmt.VarInitPrintln(`const (
	x, y = 100, 100
	a, b = 1 << iota, iota * 2 //【iota:1】 【1<<iota:2】 【iota*2:2】
	c, d                       //【iota:2】 【1<<iota:4】 【iota*2:4】
	e, f = iota, iota + 2      //【iota:3】 【iota:3】 【iota+2:5】
	g, h                       //【iota:4】 【iota:4】 【iota+2:6】
)`)
		myfmt.KeyValuePrintln("x", x, "y", y, "a", a, "b", b, "c", c, "d", d, "e", e, "f", f, "g", g, "h", h)

	})
	// go test -v -run TestName_2024_01_10_16_36_44/iota_4 github.com/gainovel/testcase/stdlib/runtime/iota
	t.Run("iota 4", func(t *testing.T) {
		// 可以用在单行const声明中(没有太大意义)
		const a = iota + 3
		fmt.Println(a)
	})
}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd docs/tests/stdlib/runtime/iota.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/runtime/iota/features_usages_001_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/iota_1 -f Makefiles/stdlib/runtime/iota.mk
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*  const (
//*          a, b = 1 << iota, iota * 2 //【iota:0】 【1<<iota:1】 【iota*2:0】
//*          c, d                       //【iota:1】 【1<<iota:2】 【iota*2:2】
//*          e, f                       //【iota:2】 【1<<iota:4】 【iota*2:4】
//*          g, h                       //【iota:3】 【1<<iota:8】 【iota*2:6】
//*          i, j                       //【iota:4】 【1<<iota:16】 【iota*2:8】
//*          k, l                       //【iota:5】 【1<<iota:32】 【iota*2:10】
//*  )
//*  --------------------------------------------------------------------------------
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                 a | 1
//*  |                 b | 0
//*  |                 c | 2
//*  |                 d | 2
//*  |                 e | 4
//*  |                 f | 4
//*  |                 g | 8
//*  |                 h | 6
//*  |                 i | 16
//*  |                 j | 8
//*  |                 k | 32
//*  |                 l | 10
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/iota_2 -f Makefiles/stdlib/runtime/iota.mk
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*  const (
//*          a, b = 1 << iota, iota * 2 //【iota:0】 【1<<iota:1】 【iota*2:0】
//*          c, d                       //【iota:1】 【1<<iota:2】 【iota*2:2】
//*          e, f = iota, iota + 2      //【iota:2】 【iota:2】 【iota+2:4】
//*          g, h                       //【iota:3】 【iota:3】 【iota+2:5】
//*          i, j = iota * 3, iota + 4  //【iota:4】 【iota*3:12】 【iota+4:8】
//*          k, l                       //【iota:5】 【iota*3:15】 【iota+4:9】
//*  )
//*  --------------------------------------------------------------------------------
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                 a | 1
//*  |                 b | 0
//*  |                 c | 2
//*  |                 d | 2
//*  |                 e | 2
//*  |                 f | 4
//*  |                 g | 3
//*  |                 h | 5
//*  |                 i | 12
//*  |                 j | 8
//*  |                 k | 15
//*  |                 l | 9
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/iota_3 -f Makefiles/stdlib/runtime/iota.mk
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*  const (
//*          x, y = 100, 100
//*          a, b = 1 << iota, iota * 2 //【iota:1】 【1<<iota:2】 【iota*2:2】
//*          c, d                       //【iota:2】 【1<<iota:4】 【iota*2:4】
//*          e, f = iota, iota + 2      //【iota:3】 【iota:3】 【iota+2:5】
//*          g, h                       //【iota:4】 【iota:4】 【iota+2:6】
//*  )
//*  --------------------------------------------------------------------------------
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                 x | 100
//*  |                 y | 100
//*  |                 a | 2
//*  |                 b | 2
//*  |                 c | 4
//*  |                 d | 4
//*  |                 e | 3
//*  |                 f | 5
//*  |                 g | 4
//*  |                 h | 6
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/iota_4 -f Makefiles/stdlib/runtime/iota.mk
//*  Test Result: 3
//**************************************************************************************
//
// qmemoutputend
