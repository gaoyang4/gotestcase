# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-11 13:10:12 星期一
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/stdlib/runtime/slice.mk
# @Version      : v0.1.0
# @Description  : 开发中···

all: 001/all

# features_usages_001_test.go
001/all: 001/share_array 001/slice_copy 001/expansion_experiment
001/index_out_of_range:
	go test -v -run TestName_2024_01_09_18_10_28/index_out_of_range github.com/gainovel/testcase/stdlib/runtime/slice
001/share_array:
	go test -v -run TestName_2024_01_09_18_10_28/share_array github.com/gainovel/testcase/stdlib/runtime/slice
001/slice_copy:
	go test -v -run TestName_2024_01_09_18_10_28/slice_copy github.com/gainovel/testcase/stdlib/runtime/slice
001/expansion_experiment:
	go test -v -run TestName_2024_01_09_18_10_28/expansion_experiment github.com/gainovel/testcase/stdlib/runtime/slice
