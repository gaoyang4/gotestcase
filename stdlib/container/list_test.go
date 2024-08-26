/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-07-24 15:45:53 星期三
 * @ProductName  : GoLand
 * @ProjectName  : testcase
 * @File         : stdlib/container/list_test.go
 * @Description  : 开发中···
 **/

package container

import (
	"container/list"
	"fmt"
	"testing"
)

func TestName_2024_07_24_15_45_53(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		type struct1 struct {
			a int
			b int
		}
		var (
			l1 list.List
		)
		l1.PushBack(&struct1{
			a: 1,
			b: 2,
		})
		l1.PushBack(&struct1{
			a: 3,
			b: 4,
		})
		l1.PushBack(&struct1{
			a: 5,
			b: 6,
		})
		fmt.Println(l1)
	})
}
