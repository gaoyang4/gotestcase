/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-02-20 12:18:40 星期二
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/struct/features_usages_001_test.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/
// qmemcodestart
package struct000

import (
	"fmt"
	"testing"
)

type Aface interface {
	Name() string
	SetName(string)
}

type A struct {
	name string
}

func (a A) Name() string {
	a.name = "Hi! " + a.name
	return a.name
}

func (a A) SetName(name string) {
	a.name = name
}
func TestName_2024_02_20_12_18_40(t *testing.T) {
	// 方法即method，Go语言支持为自定义类型实现方法，method在具体实现上与普通的函数并无不同，只不过会通过运行时栈多传递一个隐含的参数，这个隐含的参数就是所谓的接收者。
	// 以下代码展示了两种不同的写法，都能顺利通过编译并正常运行，实际上这两种写法会生成同样的机器码。
	// 第一种：a.Name()，这是我们惯用的写法，很方便；
	// 第二种：A.Name(a)，这种写法更底层也更严谨，要求所有的类型必须严格对应，否则是无法通过编译的。
	// 其实编译器会帮我们把第一种转换为第二种的形式，所以我们惯用的第一种写法只是“语法糖”，方便而已。
	// go test -v -run TestName_2024_02_20_12_18_40/case1 github.com/gainovel/testcase/stdlib/struct
	t.Run("case1", func(t *testing.T) {
		a := A{name: "eggo"}
		// 1）编译器的语法糖，提供面向对象的语法
		fmt.Println(a.Name())
		// 2）更贴近真实实现的写法，和普通函数调用几乎没什么不同
		fmt.Println(A.Name(a))
	})
	t.Run("method", func(t *testing.T) {
		var (
			a1face Aface
			a1     *A
			a2     A
		)
		a1 = &A{
			name: "xiaoming",
		}
		a2 = A{
			name: "xiaoming",
		}
		a1.SetName("xiaohong")
		a2.SetName("xiaohong")
		fmt.Println(a1.name)
		fmt.Println(a2.name)
		a1face = a2
		fmt.Printf("%T\n", a1face)
		a1face = a1
		fmt.Printf("%T\n", a1face)
		fmt.Println(a1face.Name())
	})
}

// qmemcodeend

// qmemoutputstart
//**************************************************************************************
//* 测试用例运行方法：
//*  0.前置要求：安装go（version>=1.20），安装make（可选）
//*  1.git clone git@gitee.com:gainovel/go-test-case.git 👉 cd go-test-case
//*    👉 cd docs/tests/stdlib/struct.md
//*    👉 找到对应的命令依次复制执行即可（在根目录（go-test-case）下执行命令）
//*  2.或者直接打开测试文件 stdlib/struct/features_usages_001_test.go，
//*    每个子测试上都有对应的命令，直接执行即可（在根目录（go-test-case）下执行命令）
//*  3. ❗ 注意：所有命令都在根目录下执行
//**************************************************************************************
//
//🏳️‍🌈
//**************************************************************************************
//*  Test Command: make 001/case1 -f  Makefiles/stdlib/struct.mk
//*  Test Result:
//*  Hi! eggo
//*  Hi! eggo
//**************************************************************************************
//
// qmemoutputend
