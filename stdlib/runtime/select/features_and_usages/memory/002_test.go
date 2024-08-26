/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-14 14:12:09 星期四
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/select/features_and_usages/memory/002_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package memory

import (
	"errors"
	"testing"
	"time"

	"github.com/gainovel/testcase/stdlib/runtime/select/tools"
)

// TestName_2024_03_14_14_12_09
/**
 * @author: Administrator
 * @description: select使用举例，测试命令已贴在相应位置，在terminal中打开项目根目录执行即可
 * @date: 2024-03-14 14:12:26
 */
func TestName_2024_03_14_14_12_09(t *testing.T) {
	// select语句中不包含case语句和default语句，那么协程将陷入永久性阻塞
	// go test -v -run TestName_2024_03_14_14_12_09/block_forever github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("block forever", func(t *testing.T) {
		go tools.SleepGoroutineNMinute(30)
		myfmt.ColorDescPrintln("use select to block forever...")
		select {}
	})
	// 有时我们会使用管道来传输错误，此时就可以使用select语句快速检查管道中是否有错误
	// err的另一种处理方式，不借助返回值而是借助管道
	// go test -v -run TestName_2024_03_14_14_12_09/fail_fast github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("fail fast", func(t *testing.T) {
		var (
			ch1 chan error
		)
		ch1 = make(chan error, 1)
		go func() {
			time.Sleep(time.Second * 3) // 模拟业务处理时间
			ch1 <- errors.New("fail fast error")
		}()
		select {
		case err, ok := <-ch1:
			myfmt.ColorDescPrintln("after 3s, entering case err,ok :=<-ch1")
			myfmt.KeyValuePrintln("len(ch1)", len(ch1), "cap(ch1)", cap(ch1), "err", err, "ok", ok)
		}
	})
	// err的另一种处理方式，不借助返回值而是借助管道
	// go test -v -run TestName_2024_03_14_14_12_09/wait_timeout github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
	t.Run("wait timeout", func(t *testing.T) {
		var (
			stopCh chan struct{}
			timer  *time.Timer
		)
		stopCh = make(chan struct{}, 1)
		timer = time.NewTimer(time.Second * 3)
		for {
			select {
			case <-stopCh:
				myfmt.ColorDescPrintln("entering case <-stopCh")
			case <-timer.C:
				myfmt.ColorDescPrintln("timeout, entering case <-timer.C")
			}
			timer.Reset(time.Second * 3)
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
//*  2.或者直接打开测试文件 stdlib/runtime/select/features_and_usages/memory/002_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_14_14_12_09/block_forever -f Makefiles/stdlib/runtime/select.mk
//*  Test Result:
//*  👇
//*  use select to block forever...
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: TestName_2024_03_14_14_12_09/fail_fast -f Makefiles/stdlib/runtime/select.mk
//*  Test Result:
//*  👇
//*  after 3s, entering case err,ok :=<-ch1
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |          len(ch1) | 0
//*  |          cap(ch1) | 1
//*  |               err | fail fast error
//*  |                ok | true
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_14_14_12_09/wait_timeout -f Makefiles/stdlib/runtime/select.mk
//*  Test Result:
//*  👇
//*  timeout, entering case <-timer.C
//*  👇
//*  timeout, entering case <-timer.C
//*  👇
//*  timeout, entering case <-timer.C
//**************************************************************************************
//
// qmemoutputend
