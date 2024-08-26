/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-14 22:45:35 星期四
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/for_range/features_and_usages/memory/002_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package memory

import (
	"fmt"
	"testing"
)

func TestName_2024_03_14_22_45_35(t *testing.T) {
	// 编译器会将for-range语句处理成传统的for循环，
	// 因此要遍历的元素范围在for-range转换成传统的for循环后就确定了，无法遍历到范围外的元素
	// go test -v -run TestName_2024_03_14_22_45_35/add_elem_when_range github.com/gainovel/testcase/stdlib/runtime/for_range/features_and_usages/memory/
	t.Run("add elem when range", func(t *testing.T) {
		var (
			nums1 []int
		)
		nums1 = []int{1, 2, 3}
		myfmt.VarInitPrintln(`		var (
			nums1 []int
		)
		nums1 = []int{1, 2, 3}`)
		myfmt.ColorDescPrintln("for i, v := range nums1")
		for i, v := range nums1 {
			nums1 = append(nums1, 100)
			myfmt.ColorDescPrintln(fmt.Sprintf("第%d次迭代", i+1))
			myfmt.KeyValuePrintln("i", i, "v", v, "nums1", nums1)
		}
	})
	// go test -v -run TestName_2024_03_14_22_45_35/change_elem_when_range github.com/gainovel/testcase/stdlib/runtime/for_range/features_and_usages/memory
	t.Run("change elem when range", func(t *testing.T) {
		numbers2 := [...]int{1, 2, 3, 4, 5, 6}
		maxIndex2 := len(numbers2) - 1
		for i, e := range numbers2 {
			if i == maxIndex2 {
				numbers2[0] += e
			} else {
				numbers2[i+1] += e
			}
			myfmt.KeyValuePrintln("numbers2", numbers2)
		}
		myfmt.ColorDescPrintln("finally value:")
		myfmt.KeyValuePrintln("numbers2", numbers2)
	})
}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd Makefiles/stdlib/runtime/for_range.mk
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/runtime/for_range/features_and_usages/memory/002_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_14_17_03_07/add_elem_when_range -f Makefiles/stdlib/runtime/for_range.mk
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*                  var (
//*                          nums1 []int
//*                  )
//*                  nums1 = []int{1, 2, 3}
//*  --------------------------------------------------------------------------------
//*  👇
//*  for i, v := range nums1
//*  👇
//*  第1次迭代
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                 i | 0
//*  |                 v | 1
//*  |             nums1 | [1 2 3 100]
//*  --------------------------------------------------------------------------------
//*  👇
//*  第2次迭代
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                 i | 1
//*  |                 v | 2
//*  |             nums1 | [1 2 3 100 100]
//*  --------------------------------------------------------------------------------
//*  👇
//*  第3次迭代
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |                 i | 2
//*  |                 v | 3
//*  |             nums1 | [1 2 3 100 100 100]
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
// qmemoutputend
