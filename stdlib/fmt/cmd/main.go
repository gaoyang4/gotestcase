/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-12 16:02:11 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/fmt/cmd/main.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		err  error
		dir1 string
	)
	dir1 = "D:\\Workspace\\GoLand\\gogs\\diagram\\flowchar\\data"
	executable, err := os.Executable()
	fmt.Println(executable, err, dir1)
}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd docs/tests/xxx/xxx/xxx.md
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
