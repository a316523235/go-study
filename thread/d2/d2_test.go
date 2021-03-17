package d2

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
池程池探索
*/

type job struct {
	id      int
	randNum int
}

type result struct {
	job *job
	sum int
}

func creatPool(num int, jobChan chan *job, resultChan chan *result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *job, resultChan chan *result) {
			for job := range jobChan {
				r_num := job.randNum
				sum := 0
				for r_num != 0 {
					sum += r_num % 10
					r_num /= 10
				}
				// 想要的结果是Result
				r := &result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r

				time.Sleep(1 * time.Second)
			}
		}(jobChan, resultChan)
	}
}

func TestXX(t *testing.T) {
	jobChan, resultChan := make(chan *job), make(chan *result)
	creatPool(10, jobChan, resultChan)
	go func() {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.id,
				result.job.randNum, result.sum)
		}
	}()

	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &job{
			id:      id,
			randNum: r_num,
		}
		jobChan <- job
		if id > 10 {
			break
		}
	}
}
