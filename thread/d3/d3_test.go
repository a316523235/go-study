package d3

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T)  {
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <- timer1.C
	fmt.Printf("t1:%v\n", t2)
}

func Test2(t *testing.T) {
	// 2.验证timer只能响应1次, 并在下一轮循环中报错
	timer2 := time.NewTimer(time.Second)
	for t := range timer2.C{
		fmt.Println("时间到", t)
	}
}

func Test3(t *testing.T) {
	// 2.验证timer只能响应1次, 并在下一轮循环中报错
	timer2 := time.After(time.Second)
	for t := range timer2 {
		fmt.Println("时间到", t)
	}
}

func Test4(t *testing.T) {
	x1 := make(chan<- int)
	close(x1)
}

func Test5(t *testing.T) {
	// 5.重置定时器
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
}

func Test6(t *testing.T) {
	ticker := time.NewTicker(3 * time.Second)
	for x1 := range ticker.C {
		fmt.Println(x1)
		ticker.Stop()
	}
}

func Test7(t *testing.T)  {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			t1 := <-ticker.C
			fmt.Println(t1)
			if i == 5 {
				//停止
				ticker.Stop()
			}
		}
	}()
	for {
		fmt.Println("...")
		time.Sleep(3 * time.Second)
	}
}