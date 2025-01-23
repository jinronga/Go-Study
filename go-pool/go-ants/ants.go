package go_ants

import (
	"fmt"
	"sync"
	"time"
)

// 任务
func sendMail(i int, wg *sync.WaitGroup) func() {
	var cnt int
	return func() {
		for {
			time.Sleep(time.Second * 2)
			fmt.Println("send mail to ", i)
			cnt++
			if cnt > 5 && i == 1 {
				fmt.Println("退出协程ID:", i)
				break
			}
		}
		wg.Done()
	}
}
