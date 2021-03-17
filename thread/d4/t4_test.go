package d4

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	//lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10000 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	//lock.Unlock()                     // 解互斥锁
	fmt.Println("write over")
	wg.Done()
}

func read() {
	time.Sleep(5 * time.Millisecond) // 假设读操作耗时1毫秒
	//lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	//lock.Unlock()                // 解互斥锁
	fmt.Println("read over")
	wg.Done()
}


func Test1(t *testing.T) {
	start := time.Now()

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}