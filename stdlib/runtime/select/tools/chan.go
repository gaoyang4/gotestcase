/**
 * @Author       : gainovel
 * @Organization : Copyright © 2023-2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-03-13 10:46:27 星期三
 * @ProductName  : GoLand
 * @PrjectName   : test-case
 * @File         : stdlib/runtime/select/tools/chan.go
 * @Version      : v0.1.0
 * @Description  : 开发中···
 **/

// qmemcodestart
package tools

import "time"

func SleepGoroutineNMinute(n time.Duration) {
	time.Sleep(time.Minute * n)
}
func SleepGoroutineNSeconds(n time.Duration) {
	time.Sleep(time.Second * n)
}
func SleepGoroutineNHours(n time.Duration) {
	time.Sleep(time.Hour * n)
}

func SendToIntChanNTimes(ch chan int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}
func SendToByteChan(ch chan byte) {
	var (
		str1 string
	)
	str1 = "hello world!"
	for i := 0; i < len(str1); i++ {
		ch <- str1[i]
	}
}
