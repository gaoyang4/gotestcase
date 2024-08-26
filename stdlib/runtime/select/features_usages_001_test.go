/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-12 18:57:43 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/select/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package select000

import (
	"fmt"
	"testing"
)

func sendToIntChan(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func TestName_2024_03_12_18_57_43(t *testing.T) {
	// go test -v -run TestName_2024_03_12_18_57_43/select_chan github.com/gainovel/testcase/stdlib/runtime/select
	t.Run("select chan", func(t *testing.T) {
		var (
			ch1 chan int
		)
		ch1 = make(chan int, 1)
		go sendToIntChan(ch1)
	labelFor:
		for {
			select {
			case num, ok := <-ch1:
				fmt.Println(num, ok)
				if num == 9 {
					break labelFor
				}
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
//*  2.或者直接打开测试文件 stdlib/runtime/select/features_usages_001_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_12_18_57_43/select_chan -f Makefiles/stdlib/runtime/select.mk
//*  Test Result:
//*  0 true
//*  1 true
//*  2 true
//*  3 true
//*  4 true
//*  5 true
//*  6 true
//*  7 true
//*  8 true
//*  9 true
//**************************************************************************************
//
// qmemoutputend
