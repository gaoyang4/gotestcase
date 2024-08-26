# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-11 14:06:19 星期一
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/stdlib/runtime/string.mk
# @Version      : v0.1.0
# @Description  : 开发中···

all: 001/all

# stdlib/runtime/string/features_usages_001_test.go
001/all: 001/string_to_bytes_or_runes
001/string_to_bytes_or_runes:
	go test -v -run TestName_2024_01_10_17_38_02/string_to_bytes_or_runes github.com/gainovel/testcase/stdlib/runtime/string

