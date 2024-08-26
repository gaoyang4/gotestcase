# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-11 14:27:20 星期一
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/stdlib/struct.mk
# @Version      : v0.1.0
# @Description  : 开发中···

all: 001/all

# stdlib/struct/features_usages_001_test.go
001/all: 001/case1
001/case1:
	go test -v -run TestName_2024_02_20_12_18_40/case1 github.com/gainovel/testcase/stdlib/struct

