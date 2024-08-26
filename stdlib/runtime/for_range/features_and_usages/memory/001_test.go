/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-14 17:03:07 星期四
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/for_range/features_and_usages/memory/001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package memory

import (
	"fmt"
	"testing"

	commontools "github.com/gainovel/testcase/tools/common"
)

// TestName_2024_03_14_17_03_07
/**
 * @author: Administrator
 * @description: for-range用法测试，测试命令已贴在相应位置，在terminal中打开项目根目录执行即可
 * @date: 2024-03-14 19:25:36
 */
func TestName_2024_03_14_17_03_07(t *testing.T) {
	// for-range表达式用于遍历所有的集合类型，包括数组、切片、string、map、chan
	// for range返回值的个数。除了chan最多返回一个值外，其它的集合类型最多返回两个值
	// go test -v -run TestName_2024_03_14_17_03_07/case1 github.com/gainovel/testcase/stdlib/runtime/for_range/features_and_usages/memory/
	t.Run("case1", func(t *testing.T) {
		var (
			arr1 [3]int
			s1   []int
			str1 string
			map1 map[int]int
			ch1  chan int
		)
		arr1 = [3]int{1, 2, 3}
		s1 = []int{4, 5, 6}
		str1 = "hello world"
		map1 = make(map[int]int)
		map1[1] = 2
		ch1 = make(chan int, 2)
		ch1 <- 8
		ch1 <- 9
		myfmt.VarInitPrintln(`		var (
			arr1 [3]int
			s1   []int
			str1 string
			map1 map[int]int
			ch1  chan int
		)
		arr1 = [3]int{1, 2, 3}
		s1 = []int{4, 5, 6}
		str1 = "hello world"
		map1 = make(map[int]int)
		map1[1] = 2
		ch1 = make(chan int, 2)
		ch1 <- 8
		ch1 <- 9`)
		// for-range 遍历数组，一个返回值
		myfmt.ColorDescPrintln("for-range 遍历数组，一个返回值")
		for i := range arr1 {
			fmt.Printf("i:%d\t", i)
		}
		fmt.Println()
		// for-range 遍历数组，两个返回值
		myfmt.ColorDescPrintln("for-range 遍历数组，两个返回值")
		for i, v := range arr1 {
			fmt.Printf("i:%d,v:%d\t", i, v)
		}
		fmt.Println()
		fmt.Println(commontools.GenerateHorizontalLine(80))

		// for-range 遍历切片，一个返回值
		myfmt.ColorDescPrintln("for-range 遍历切片，一个返回值")
		for i := range s1 {
			fmt.Printf("i:%d\t", i)
		}
		fmt.Println()
		// for-range 遍历切片，两个返回值
		myfmt.ColorDescPrintln("for-range 遍历切片，两个返回值")
		for i, v := range s1 {
			fmt.Printf("i:%d,v:%d\t", i, v)
		}
		fmt.Println()
		fmt.Println(commontools.GenerateHorizontalLine(80))

		// for-range 遍历字符串，一个返回值
		myfmt.ColorDescPrintln("for-range 遍历字符串，一个返回值")
		for i := range str1 {
			fmt.Printf("i:%d\t", i)
		}
		fmt.Println()
		// for-range 遍历字符串，两个返回值
		myfmt.ColorDescPrintln("for-range 遍历字符串，两个返回值")
		for i, v := range str1 {
			fmt.Printf("i:%d,v:%d\t", i, v)
		}
		fmt.Println()
		fmt.Println(commontools.GenerateHorizontalLine(80))

		// for-range 遍历map，一个返回值
		myfmt.ColorDescPrintln("for-range 遍历map，一个返回值")
		for key := range map1 {
			fmt.Printf("key:%d\t", key)
		}
		fmt.Println()
		// for-range 遍历map，两个返回值
		myfmt.ColorDescPrintln("for-range 遍历map，两个返回值")
		for key, val := range map1 {
			fmt.Printf("key:%d,val:%d\t", key, val)
		}
		fmt.Println()
		fmt.Println(commontools.GenerateHorizontalLine(80))
		close(ch1)
		// for-range 遍历chan，一个返回值
		myfmt.ColorDescPrintln("for-range 遍历chan，一个返回值")
		for val := range ch1 {
			fmt.Printf("val:%d\t", val)
		}
		fmt.Println()
		fmt.Println(commontools.GenerateHorizontalLine(80))
	})
	// 迭代变量共享。在for-range表达式中，迭代变量只会声明一次，在多次迭代中共享。
	// go test -v -run TestName_2024_03_14_17_03_07/share github.com/gainovel/testcase/stdlib/runtime/for_range/features_and_usages/memory/
	t.Run("share", func(t *testing.T) {
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
			myfmt.ColorDescPrintln(fmt.Sprintf("第%d次迭代", i+1))
			myfmt.KeyValuePrintln("i address", &i, "i", i, "v address", &v, "v", v)
		}
	})
	// for range作用于string时，是按照unicode码点进行遍历的，而不是逐个byte进行遍历
	// go test -v -run TestName_2024_03_14_17_03_07/string_to_bytes_or_runes github.com/gainovel/testcase/stdlib/runtime/for_range/features_and_usages/memory/
	t.Run("string to bytes or runes", func(t *testing.T) {
		var (
			str1  string
			bytes []byte
			runes []rune
		)
		str1 = "hello 中国！"
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// ['h','e','l','l','o',   ' ',         '中',          '国',           '！'   ]
		// [104,101,108,108,111,    32,     228,184,173    229,155,189    239,188,129] []byte len:15
		// [104,101,108,108,111,    32,        20013,        22269,          65281   ] []rune len:9
		myfmt.VarInitPrintln(`var (
	str1  string
	bytes []byte
	runes []rune
)`)
		bytes = []byte(str1)
		runes = []rune(str1)
		myfmt.ColorDescPrintln("bytes = []byte(str1)", "runes = []rune(str1)")
		myfmt.KeyValuePrintln("str1", str1, "bytes", bytes, "runes", runes)
	})

}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd docs/tests/stdlib/runtime/for_range.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/runtime/for_range/features_and_usages/memory/001_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_14_17_03_07/case1 -f Makefiles/stdlib/runtime/for_range.mk
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*                  var (
//*                          arr1 [3]int
//*                          s1   []int
//*                          str1 string
//*                          map1 map[int]int
//*                          ch1  chan int
//*                  )
//*                  arr1 = [3]int{1, 2, 3}
//*                  s1 = []int{4, 5, 6}
//*                  str1 = "hello world"
//*                  map1 = make(map[int]int)
//*                  map1[1] = 2
//*                  ch1 = make(chan int, 2)
//*                  ch1 <- 8
//*                  ch1 <- 9
//*  --------------------------------------------------------------------------------
//*  👇
//*  for-range 遍历数组，一个返回值
//*  i:0     i:1     i:2
//*  👇
//*  for-range 遍历数组，两个返回值
//*  i:0,v:1 i:1,v:2 i:2,v:3
//*  --------------------------------------------------------------------------------
//*  👇
//*  for-range 遍历切片，一个返回值
//*  i:0     i:1     i:2
//*  👇
//*  for-range 遍历切片，两个返回值
//*  i:0,v:4 i:1,v:5 i:2,v:6
//*  --------------------------------------------------------------------------------
//*  👇
//*  for-range 遍历字符串，一个返回值
//*  i:0     i:1     i:2     i:3     i:4     i:5     i:6     i:7     i:8     i:9     i:10
//*  👇
//*  for-range 遍历字符串，两个返回值
//*  i:0,v:104       i:1,v:101       i:2,v:108       i:3,v:108       i:4,v:111       i:5,v:32        i:6,v:119       i:7,v:111       i:8,v:114       i:9,v:108       i:10,v:100
//*  --------------------------------------------------------------------------------
//*  👇
//*  for-range 遍历map，一个返回值
//*  key:1
//*  👇
//*  for-range 遍历map，两个返回值
//*  key:1,val:2
//*  --------------------------------------------------------------------------------
//*  👇
//*  for-range 遍历chan，一个返回值
//*  val:8   val:9
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_14_17_03_07/share -f Makefiles/stdlib/runtime/for_range.mk
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
//*  |         i address | 0xc00009eba0
//*  |                 i | 0
//*  |         v address | 0xc00009eba8
//*  |                 v | 1
//*  --------------------------------------------------------------------------------
//*  👇
//*  第2次迭代
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |         i address | 0xc00009eba0
//*  |                 i | 1
//*  |         v address | 0xc00009eba8
//*  |                 v | 2
//*  --------------------------------------------------------------------------------
//*  👇
//*  第3次迭代
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |         i address | 0xc00009eba0
//*  |                 i | 2
//*  |         v address | 0xc00009eba8
//*  |                 v | 3
//*  --------------------------------------------------------------------------------*
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make TestName_2024_03_14_17_03_07/string_to_bytes_or_runes -f Makefiles/stdlib/runtime/for_range.mk
//*  Test Result:
//*  👇
//*  变量初始化：
//*  --------------------------------------------------------------------------------
//*  var (
//*          str1  string
//*          bytes []byte
//*          runes []rune
//*  )
//*  --------------------------------------------------------------------------------
//*  👇
//*  bytes = []byte(str1)
//*  runes = []rune(str1)
//*  --------------------------------------------------------------------------------
//*  |               key | value
//*  --------------------------------------------------------------------------------
//*  |              str1 | hello 中国！
//*  |             bytes | [104 101 108 108 111 32 228 184 173 229 155 189 239 188 129]
//*  |             runes | [104 101 108 108 111 32 20013 22269 65281]
//*  --------------------------------------------------------------------------------
//**************************************************************************************
//
// qmemoutputend
