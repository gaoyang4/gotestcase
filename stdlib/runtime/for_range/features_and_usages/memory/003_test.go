/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-23 17:11:54 星期六
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/for_range/features_and_usages/memory/003_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package memory

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("setup")
	m.Run()
	println("tear down")
	os.Exit(-1)
}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd testdoc/xxx/xxx/xxx.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 xxx/xxx/xxx.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: xxx
//*  Test Result:
//*
//**************************************************************************************
//
// qmemoutputend
