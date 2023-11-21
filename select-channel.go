package main

import (
	"fmt"
	"time"
)

/*
select 是 Go 中的一个控制结构，类似于 switch 语句。
select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。

由于 select 语句的特性，break 语句并不能直接用于跳出 select 语句本身，
因为 select 语句是非阻塞的，它会一直等待所有的通信操作都准备就绪。
如果需要提前结束 select 语句的执行，可以使用 return 或者 goto 语句来达到相同的效果。
*/
func process(ch chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Println("Received value:", val)
			// 执行一些逻辑
			if val == 5 {
				return // 提前结束 select 语句的执行
			}
		default:
			fmt.Println("No value received yet.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ch := make(chan int)

	go process(ch)

	time.Sleep(2 * time.Second)
	ch <- 1
	time.Sleep(1 * time.Second)
	ch <- 3
	time.Sleep(1 * time.Second)
	ch <- 5
	time.Sleep(1 * time.Second)
	ch <- 7

	time.Sleep(2 * time.Second)
}

/*
No value received yet.
No value received yet.
No value received yet.
No value received yet.
Received value: 1
No value received yet.
No value received yet.
No value received yet.
Received value: 3
No value received yet.
No value received yet.
Received value: 5
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        D:/Develop/GoConcurrents/go-select-channel.go:46 +0xe5
exit status 2
*/