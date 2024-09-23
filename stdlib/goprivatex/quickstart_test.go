/**
 * @Author       : GaoYang
 * @Date         : 2024-09-23 22:24:32 星期一
 * @Organization : Copyright © 2024 gainovel.com All Rights Reserved.
 * @ProductName  : GoLand
 * @ProjectName  : testcase
 * @File         : stdlib/goprivatex/quickstart_test.go
 * @Description  : ...
 **/

package goprivatex

import (
	"fmt"
	"gaoyang4.shenzhuo.vip/goget/go/privaterepo"
	"gaoyang4.shenzhuo.vip/goget/go/privaterepo/lib01"
	"testing"
)

func TestName_2024_09_23_22_24_32(t *testing.T) {
	t.Run("case0", func(t *testing.T) {
		fmt.Println("go private use:")
		lib00Name := privaterepo.Version()
		fmt.Println(lib00Name)
	})
	t.Run("case1", func(t *testing.T) {
		fmt.Println("go private use:")
		lib00Name := lib01.Name()
		fmt.Println(lib00Name)
	})
}
