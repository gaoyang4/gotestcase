/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-15 10:59:22 星期五
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/fmt/features_usages/memory/001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package memory

import (
	"fmt"
	"math"
	"runtime"
	"testing"
	"time"
)

func TestName_2024_03_15_10_59_22(t *testing.T) {
	// go test -v count 2 -run TestName_2024_03_15_10_59_22/hello
	t.Run("hello", func(t *testing.T) {
		fmt.Println("hello world!")
	})
	// go test -v -run TestName_2024_03_15_10_59_22/case1 github.com/gainovel/testcase/stdlib/fmt/features_usages/memory
	t.Run("case1", func(t *testing.T) {
		var (
			coordinateSystem [][]rune
			n                int
			x                []float64
			y                []float64
		)
		n = 32
		coordinateSystem = make([][]rune, n)
		for i := 0; i < n; i++ {
			coordinateSystem[i] = make([]rune, n)
		}
		coordinateSystem[13][0] = '消'
		coordinateSystem[14][0] = '费'
		coordinateSystem[15][0] = '能'
		coordinateSystem[16][0] = '力'

		//for i := len(coordinateSystem) - 1; i >= 0; i-- {
		//	fmt.Println(string(coordinateSystem[i]))
		//}
		for i := 1; i <= 80; i++ {
			x = append(x, float64(i))
			y = append(y, -1/32*(float64(i)-32)*(float64(i)-32)+32.0)
		}
		fmt.Println(x, y)
	})
	// 生成y = -1/32*(x-32)²+32 坐标轴数据
	t.Run("case2", func(t *testing.T) {
		var (
			x_axis []string
			y_axis []string
		)
		for i := 1; i < 100; i++ {
			x_axis = append(x_axis, fmt.Sprintf("%d,", i))
			y_axis = append(y_axis, fmt.Sprintf("%.2f", -1.0/32.0*math.Pow(float64(i)-32.0, 2.0)+32.0))
		}
		fmt.Println(x_axis)
		fmt.Println(y_axis)
	})
	t.Run("", func(t *testing.T) {
		runtime.GOMAXPROCS(1)
		go func() {
			for {

			}
		}()
		time.Sleep(1 * time.Second)
		println("done")
	})
}

func TestName(t *testing.T) {

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
