# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-11 14:15:59 星期一
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/stdlib/runtime/iota.mk
# @Version      : v0.1.0
# @Description  : 开发中···

all: 001/all

# stdlib/runtime/iota/features_usages_001_test.go
001/all: 001/index_out_of_range
001/iota_1:
	go test -v -run TestName_2024_01_10_16_36_44/iota_1 github.com/gainovel/testcase/stdlib/runtime/iota
001/iota_2:
	go test -v -run TestName_2024_01_10_16_36_44/iota_2 github.com/gainovel/testcase/stdlib/runtime/iota
001/iota_3:
	go test -v -run TestName_2024_01_10_16_36_44/iota_3 github.com/gainovel/testcase/stdlib/runtime/iota
001/iota_4:
	go test -v -run TestName_2024_01_10_16_36_44/iota_4 github.com/gainovel/testcase/stdlib/runtime/iota

