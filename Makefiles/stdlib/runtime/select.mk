# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-12 19:02:14 星期二
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/stdlib/runtime/select.mk
# @Version      : v0.1.0
# @Description  : 开发中···

# stdlib/runtime/select/features_usages_001_test.go
TestName_2024_03_12_18_57_43/all: TestName_2024_03_12_18_57_43/select_chan
TestName_2024_03_12_18_57_43/select_chan:
	go test -v -run TestName_2024_03_12_18_57_43/select_chan github.com/gainovel/testcase/stdlib/runtime/select

# stdlib/runtime/select/features_and_usages/comprehensive/001.go
TestName_2024_03_13_10_44_15/all: TestName_2024_03_13_10_44_15/case1
TestName_2024_03_13_10_44_15/case1:
	go test -v -run TestName_2024_03_13_10_44_15/case1 github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/comprehensive

# stdlib/runtime/select/features_and_usages/memory/001_test.go
TestName_2024_03_13_10_43_52/choose_one_case_to_run:
	go test -v -run TestName_2024_03_13_10_43_52/choose_one_case_to_run github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
TestName_2024_03_13_10_43_52/select_block_case1:
	go test -v -run TestName_2024_03_13_10_43_52/select_block_case1 github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
TestName_2024_03_13_10_43_52/select_block_case2:
	go test -v -run TestName_2024_03_13_10_43_52/select_block_case2 github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
TestName_2024_03_13_10_43_52/read_from_a_chan:
	go test -v -run TestName_2024_03_13_10_43_52/read_from_a_chan github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
TestName_2024_03_13_10_43_52/chan_returen_requirement:
	go test -v -run TestName_2024_03_13_10_43_52/chan_returen_requirement github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
TestName_2024_03_13_10_43_52/default:
	go test -v -run TestName_2024_03_13_10_43_52/default github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory

# stdlib/runtime/select/features_and_usages/memory/002_test.go

TestName_2024_03_14_14_12_09/block_forever:
	go test -v -run TestName_2024_03_14_14_12_09/block_forever github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
TestName_2024_03_14_14_12_09/fail_fast:
	go test -v -run TestName_2024_03_14_14_12_09/fail_fast github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
TestName_2024_03_14_14_12_09/wait_timeout:
	go test -v -run TestName_2024_03_14_14_12_09/wait_timeout github.com/gainovel/testcase/stdlib/runtime/select/features_and_usages/memory
