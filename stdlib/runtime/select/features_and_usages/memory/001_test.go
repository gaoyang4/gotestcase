/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-13 10:43:52 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/select/features_and_usages/memory/001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package memory

import (
	"fmt"
	"testing"

	"github.com/gainovel/testcase/stdlib/runtime/select/tools"
	commonprint "github.com/gainovel/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

// TestName_2024_03_13_10_43_52
/**
 * @author: Administrator
 * @description: select特性测试，测试命令已贴在相应位置，在terminal中打开项目根目录执行即可
 * @date: 2024-03-14 14:11:40
 */
func TestName_2024_03_13_10_43_52(t *testing.T) {
	// select的每个case语句只能操作一个管道，要么写入数据，要么读取数据
	// 如果多个case语句均没有阻塞，那么select将随机挑选一个case执行
	// go test -v -run TestName_2024_03_13_10_43_52/choose_one_case_to_run github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("choose one case to run", func(t *testing.T) {
		var (
			ch1 chan int
			ch2 chan int
		)
		ch1 = make(chan int, 1)
		ch2 = make(chan int, 1)
		ch1 <- 1
		ch2 <- 2
		myfmt.VarInitPrintln(`		var (
			ch1 chan int
			ch2 chan int
		)
		ch1 = make(chan int, 1)
		ch2 = make(chan int, 1)
		ch1 <- 1
		ch2 <- 2`)
		myfmt.ColorDescPrintln("ch1 status")
		myfmt.KeyValuePrintln("ch1", ch1, "len(ch1)", len(ch1), "cap(ch1)", cap(ch1))
		myfmt.ColorDescPrintln("ch2 status")
		myfmt.KeyValuePrintln("ch2", ch2, "len(ch2)", len(ch2), "cap(ch2)", cap(ch2))
		go tools.SendToIntChanNTimes(ch1, 100)
		go tools.SendToIntChanNTimes(ch2, 100)
		for i := 0; i < 8; i++ {
			myfmt.ColorDescPrintln(fmt.Sprintf("第%d次选择", i+1))
			select {
			case <-ch1:
				myfmt.ColorDescPrintln("entering case <-ch1")
			case <-ch2:
				myfmt.ColorDescPrintln("entering case <-ch2")
			}
		}
	})
	// 鉴于管道的特性，如果管道中没有数据则 读取操作 会阻塞，
	// go test -v -run TestName_2024_03_13_10_43_52/select_block_case1 github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("select block case1", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		go tools.SleepGoroutineNMinute(30)
		myfmt.VarInitPrintln(`		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		go tools.SleepGoroutineNMinute(30)`)
		myfmt.ColorDescPrintln("ch1 status")
		myfmt.KeyValuePrintln("ch1", ch1, "len(ch1)", len(ch1), "cap(ch1)", cap(ch1))
		myfmt.ColorDescPrintln("ch1 has no data, select read blocked...")
		select {
		case <-ch1:
			myfmt.ColorDescPrintln("entering case <-ch1")
		}
		myfmt.ColorDescPrintln("finish...")
	})
	// 如果管道中没有空余的缓冲区则 写入操作 会阻塞
	// go test -v -run TestName_2024_03_13_10_43_52/select_block_case2 github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("select block case2", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		ch1 <- 1
		ch1 <- 2
		go tools.SleepGoroutineNMinute(30)
		myfmt.VarInitPrintln(`		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		ch1 <- 1
		ch1 <- 2
		go tools.SleepGoroutineNMinute(30)`)
		myfmt.ColorDescPrintln("ch1 status")
		myfmt.KeyValuePrintln("ch1", ch1, "len(ch1)", len(ch1), "cap(ch1)", cap(ch1))
		myfmt.ColorDescPrintln("ch1 has no space, select write blocked...")
		select {
		case ch1 <- 3:
			myfmt.ColorDescPrintln("entering case <-ch1")
		}
		myfmt.ColorDescPrintln("finish...")
	})
	// 当select的多个case语句中的管道均阻塞时，
	// 整个select语句也会陷入阻塞（没有default语句的情况下），
	// 直到任意一个管道解除阻塞

	// 可以使用逗号 ok模式从select监听的管道中获取数据，
	// 其中ok表示是否成功地读取到了数据。
	// go test -v -run TestName_2024_03_13_10_43_52/read_from_a_chan github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("read from a chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		ch1 <- 1
		ch1 <- 2
		for i := 0; i < 2; i++ {
			select {
			case val, ok := <-ch1:
				myfmt.ColorDescPrintln("entering case val,ok := <-ch1")
				myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "val", val, "ok", ok)
			}
		}
	})
	// 管道的读操作有两个返回条件，
	// 一是成功读到数据，二是管道中已没有数据且被关闭
	// go test -v -run TestName_2024_03_13_10_43_52/chan_returen_requirement github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("chan returen requirement", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		ch1 <- 1
		ch1 <- 2
		for i := 0; i < 4; i++ {
			select {
			case val, ok := <-ch1:
				myfmt.ColorDescPrintln("entering case val,ok := <-ch1")
				myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "val", val, "ok", ok)
				if val == 2 {
					close(ch1)
				}
			}
		}
	})
	// 当select的所有语句都阻塞时，default语句将被执行
	// go test -v -run TestName_2024_03_13_10_43_52/default github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("default", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 2)
		for i := 0; i < 8; i++ {
			select {
			case <-ch1:
				myfmt.ColorDescPrintln("entering case <-ch1")
			default:
				myfmt.ColorDescPrintln("entering default")
			}
		}
	})
}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd docs/tests/stdlib/runtime/select.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/runtime/select/features_and_usages/memory/001_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: TestName_2024_03_13_10_43_52/choose_one_case_to_run
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*                  var (
//*                          ch1 chan int
//*                          ch2 chan int
//*                  )
//*                  ch1 = make(chan int, 1)
//*                  ch2 = make(chan int, 1)
//*                  ch1 <- 1
//*                  ch2 <- 2
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch1 status
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |               ch1 | 0xc000018700
//*  |          len(ch1) | 1
//*  |          cap(ch1) | 1
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch2 status
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |               ch2 | 0xc000018770
//*  |          len(ch2) | 1
//*  |          cap(ch2) | 1
//*  --------------------------------------------------------------------------------
//*  👇
//*  第1次选择
//*  👇
//*  entering case <-ch2
//*  👇
//*  第2次选择
//*  👇
//*  entering case <-ch1
//*  👇
//*  第3次选择
//*  👇
//*  entering case <-ch2
//*  👇
//*  第4次选择
//*  👇
//*  entering case <-ch1
//*  👇
//*  第5次选择
//*  👇
//*  entering case <-ch2
//*  👇
//*  第6次选择
//*  👇
//*  entering case <-ch1
//*  👇
//*  第7次选择
//*  👇
//*  entering case <-ch1
//*  👇
//*  第8次选择
//*  👇
//*  entering case <-ch2
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_13_10_43_52/select_block_case1 -f Makefiles/stdlib/runtime/select.mk
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*                  var (
//*                          ch1 chan int
//*                  )
//*                  ch1 = make(chan int, 2)
//*                  go tools.SleepGoroutineNMinute(30)
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch1 status
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |               ch1 | 0xc000018700
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch1 has no data, select read blocked...
//*
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: go test -v -run TestName_2024_03_13_10_43_52/select_block_case2 github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*                  var (
//*                          ch1 chan int
//*                  )
//*                  ch1 = make(chan int, 2)
//*                  ch1 <- 1
//*                  ch1 <- 2
//*                  go tools.SleepGoroutineNMinute(30)
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch1 status
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |               ch1 | 0xc00012a4d0
//*  |          len(ch1) | 2
//*  |          cap(ch1) | 2
//*  --------------------------------------------------------------------------------
//*  👇
//*  ch1 has no space, select write blocked...
//*
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_13_10_43_52/read_from_a_chan -f Makefiles/stdlib/runtime/select.mk
//*  Test Result:
//*  👇
//*  entering case val,ok := <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 1
//*  |          cap(ch1) | 2
//*  |               val | 1
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//*  👇
//*  entering case val,ok := <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  |               val | 2
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_13_10_43_52/chan_returen_requirement -f Makefiles/stdlib/runtime/select.mk
//*  Test Result:
//*  👇
//*  entering case val,ok := <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 1
//*  |          cap(ch1) | 2
//*  |               val | 1
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//*  👇
//*  entering case val,ok := <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  |               val | 2
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//*  👇
//*  entering case val,ok := <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  |               val | 0
//*  |                ok | false
//*  --------------------------------------------------------------------------------
//*  👇
//*  entering case val,ok := <-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 2
//*  |               val | 0
//*  |                ok | false
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_13_10_43_52/default -f Makefiles/stdlib/runtime/select.mk
//*  Test Result:
//*  👇
//*  entering default
//*  👇
//*  entering default
//*  👇
//*  entering default
//*  👇
//*  entering default
//*  👇
//*  entering default
//*  👇
//*  entering default
//*  👇
//*  entering default
//*  👇
//*  entering default
//**************************************************************************************
//
// qmemoutputend
