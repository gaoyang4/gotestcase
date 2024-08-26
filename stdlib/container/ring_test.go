/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-07-24 17:26:51 星期三
 * @ProductName  : GoLand
 * @ProjectName  : testcase
 * @File         : stdlib/container/ring_test.go
 * @Description  : 开发中···
 **/

package container

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestName_2024_07_24_17_26_51(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		r := ring.New(10)
		//给闭环中的元素附值
		for i := 1; i <= r.Len(); i++ {
			r.Value = i
			r = r.Next()
		}
		r.Do(func(a any) {
			fmt.Println(a)
		})
		r = r.Move(2)
		fmt.Println(r)
	})
}
