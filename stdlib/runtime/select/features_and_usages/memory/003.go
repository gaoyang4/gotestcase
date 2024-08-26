/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-14 14:35:46 星期四
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/select/features_and_usages/memory/003.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package memory

import (
	"testing"
)

// TestName_2024_03_14_14_35_46
/**
 * @author: Administrator
 * @description: select实现原理探索 | 基于go1.15.15
 * @date: 2024-03-14 14:35:51
 */

//数据结构
// select中的case语句对应于runtime包中的scase(select-case)数据结构
//type scase struct {
//	c           *hchan         // chan
//	elem        unsafe.Pointer // data element
//	kind        uint16
//	pc          uintptr // race pc (for race detector / msan)
//	releasetime int64
//}

func TestName_2024_03_14_14_35_46(t *testing.T) {
	t.Run("case1", func(t *testing.T) {

	})
}

// qmemcodeend

// 1qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd testdoc/xxx/xxx/xxx.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/runtime/select/features_and_usages/memory/003.go，
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
