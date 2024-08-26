/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-07-25 22:26:13 星期四
 * @ProductName  : GoLand
 * @ProjectName  : testcase
 * @File         : stdlib/runtime/func/calculate.go
 * @Description  : 开发中···
 **/

package funct

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	fmt.Println(&op)
	return op(x, y), nil
}
