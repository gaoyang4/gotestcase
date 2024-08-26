/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024/1/10 17:38:02 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/string/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/
// qmemcodestart
package string

import (
	"testing"

	commonprint "github.com/gainovel/testcase/tools/common/print"
)

var (
	myfmt = commonprint.MyFmt
)

func TestName_2024_01_10_17_38_02(t *testing.T) {
	// go test -v -run TestName_2024_01_10_17_38_02/string_to_bytes_or_runes github.com/gainovel/testcase/stdlib/runtime/string
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
//*    👉 cd docs/tests/stdlib/runtime/string.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/runtime/string/features_usages_001_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: xxx
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
