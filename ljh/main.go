package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义缓冲区
// var myCh = make(chan int, 0)
var myCh = make(chan int)
var exitCh = make(chan bool, 1)

// 定义生产者
func producer() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Println("生产者生产了: ", num)
		// 往管道中写入数据
		myCh <- num
		//time.Sleep(time.Millisecond * 500)
	}
	// 生产完毕之后关闭管道
	close(myCh)
	fmt.Println("生产者停止生产")
}

// 定义消费者
func consumer() {
	// 不断从管道中获取数据, 直到管道关闭位置
	for {
		if num, ok := <-myCh; !ok {
			break
		} else {
			fmt.Println("---消费者消费了", num)
		}
	}
	fmt.Println("消费者停止消费")
	exitCh <- true
}

func main() {
	go producer()
	go consumer()
	fmt.Println("exitCh之前代码")
	<-exitCh
	fmt.Println("exitCh之后代码")
}
