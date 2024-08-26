# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-11 13:50:10 星期一
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/stdlib/runtime/map.mk
# @Version      : v0.1.0
# @Description  : 开发中···

all: 001/all

# stdlib/runtime/map/features_usages_001_test.go
001/all: 001/concurrent_map_write
001/concurrent_map_write:
	go test -v -run TestName_2024_01_10_11_22_46/concurrent_map_write github.com/gainovel/testcase/stdlib/runtime/map
001/assignment_to_entry_in_nil_map:
	go test -v -run TestName_2024_01_10_11_22_46/assignment_to_entry_in_nil_map github.com/gainovel/testcase/stdlib/runtime/map
001/map_crud:
	go test -v -run TestName_2024_01_10_11_22_46/map_crud github.com/gainovel/testcase/stdlib/runtime/map

