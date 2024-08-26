/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-07-25 16:32:00 星期四
 * @ProductName  : GoLand
 * @ProjectName  : testcase
 * @File         : stdlib/runtime/func/type_func_test.go
 * @Description  : 开发中···
 **/

package funct

import (
	"fmt"
	"testing"
)

type Printer func(string) (int, error)

func printToStd1(contents1 string) (bytesNum1 int, err1 error) {
	return fmt.Println("printToStd1 " + contents1)
}
func printToStd2(contents2 string) (bytesNum2 int, err2 error) {
	return fmt.Println("printToStd2 " + contents2)
}
func printToStd3(contents3 string) (bytesNum3 int, err3 error) {
	return fmt.Println("printToStd3 " + contents3)
}

func TestName_2024_07_25_16_32_00(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		var p Printer
		p = printToStd1
		p("something")
		p = printToStd2
		p("something")
		p = printToStd3
		p("something")
		var (
			a int
		)
		a++
		//pointer := uintptr(unsafe.Pointer(&a))
		//fmt.Println(a)
		t.Fail()
	})
	/*
		printToStd1 something
		printToStd2 something
		printToStd3 something
	*/
}
