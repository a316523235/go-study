package d1

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	for i := 1; i < 10; i++ {
		go hello()
		fmt.Println("over")
		fmt.Printf("-----%d-----\n\r", i)
	}
}

func hello() {
	fmt.Println("say hello")
}

func Test2(t *testing.T) {
	for i := 1; i < 10; i++ {
		ch := make(chan int)
		go hello2(ch)
		<-ch
		fmt.Println("over")
		fmt.Printf("-----%d-----\n\r", i)
	}
}

func hello2(ch chan int) {
	fmt.Println("say hello2")
	ch <- 1
}

func Test3(t *testing.T) {
	ch := make(chan int)
	go func(ch chan int) {
		defer func(ch chan int) { ch <- 1 }(ch)
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}(ch)

	<-ch
}

func Test4(t *testing.T) {
	ch := make(chan int)
	go recv(ch)
	x := <-ch
	fmt.Println("over", x, ch)
}

func recv(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 10
	fmt.Println("send in")
}

func Test5(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
		fmt.Println("go1 over")
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for i := 0; i < 110; i++ {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch2)
		fmt.Println("go2 over")
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

//使用channel后， 要么关闭，要么有线程在继续处理（不论这个线程是否有用到当前channel）， 否则要抱死锁
//比例此示例，或第一个go线程中未将channel关闭， 到第二个go线程结束后就会报错
func Test6(t *testing.T) {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close(c)
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("go2 over")
	}()

	for {
		x, ok := <-c
		if !ok {
			break
		}
		fmt.Println(x)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("over")
}

//会报错， 原因main函数或TestXX函数， 本身就是一个协程的执行， 所以执行执行到 c<-1时：
// 执行Test7函数的协程将被阻塞， 根本就不会往下执行去启动go func协程。 那就变成一直阻塞在c <-1了
func Test7(t *testing.T) {
	c := make(chan int)

	c <- 1

	//根本不会执行到这里
	go func() {
		c <- 1
		c <- 2
		fmt.Println("go1")
	}()

	time.Sleep(3 * time.Second)
}

//不会报错
func Test8(t *testing.T) {
	c := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		<-c
	}()
	time.Sleep(3 * time.Second)
}

//不会报错
func Test9(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		fmt.Println("go1")
	}()

	time.Sleep(3 * time.Second)
}

func Test10(t *testing.T) {
	pipline := make(chan string)
	go func() {
		pipline <- "hello world"
		pipline <- "hello China"
		//close(pipline)
	}()

	for data := range pipline {
		fmt.Println(data)
	}
}
