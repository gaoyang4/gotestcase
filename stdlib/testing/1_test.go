/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-10 20:53:36 星期日
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/testing/1_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

package testing000

import (
	"fmt"
	"testing"
)

// qmemcodestart
func TestName_2024_03_10_20_53_36(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		fmt.Println("hello world")
	})
	t.Run("case1", func(t *testing.T) {
		fmt.Println("hello world")
	})
	t.Run("case1", func(t *testing.T) {
		fmt.Println("hello world")
	})
	t.Run("case1", func(t *testing.T) {
		fmt.Println("hello world")
	})
}

func TestName_2024_03_10_20_53_37(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		fmt.Println("hello world")
	})
	t.Run("case1", func(t *testing.T) {
		fmt.Println("hello world")
	})
	t.Run("case1", func(t *testing.T) {
		fmt.Println("hello world")
	})
	t.Run("case1", func(t *testing.T) {
		fmt.Printf(" ")
	})
	t.Parallel()
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func1 := func() string {
			var (
				str1 string
			)
			str1 = "hello" + "world"
			return str1
		}
		func1()
	}
}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}

//qmemcodeend

//qmemoutputstart
//hello world
//qmemoutputend
