# @Author       : gainovel
# @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
# @Date         : 2024-03-11 14:33:28 星期一
# @ProductName  : GoLand
# @PrjectName   : test-case
# @File         : Makefiles/stdlib/sync/map.mk
# @Version      : v0.1.0
# @Description  : 开发中···

all: 001/all

# stdlib/sync/map/features_usages_001_test.go
001/all: 001/sync.Map_1 001/sync.Map.LoadOrStore 001/sync.Map.LoadAndDelete
001/sync.Map_1:
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map_1 github.com/gainovel/testcase/stdlib/sync/map
001/sync.Map.LoadOrStore:
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadOrStore github.com/gainovel/testcase/stdlib/sync/map
001/sync.Map.LoadAndDelete:
	go test -v -run TestName_2024_01_11_16_18_49/sync.Map.LoadAndDelete github.com/gainovel/testcase/stdlib/sync/map

