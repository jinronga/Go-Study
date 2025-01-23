package playground_pool

import (
	"fmt"
	"gopkg.in/go-playground/pool.v3"
	"testing"
	"time"
)

func Test_playground_pool(t *testing.T) {
	p := pool.NewLimited(20)

	defer p.Close()

	batch := p.Batch()

	timeout := time.After(10 * time.Second)

	go func() {
		for i := 0; i < 100; i++ {
			batch.Queue(SendMail(i)) // 往批量任务中添加workFunc任务
		}

		batch.QueueComplete()
	}()

	r := batch.Results()

LOOP:
	for {
		select {
		case <-timeout:
			// 超时通知
			fmt.Println("超时通知")
			break LOOP
		case email, ok := <-r:
			// 读取结果集
			if ok {
				if err := email.Error(); err != nil {
					fmt.Println("读取结果集错误，error info:", err.Error())
				}
				fmt.Println("错误结果集:", email.Value())
			} else {
				fmt.Println("finish")
				break LOOP
			}
		}
	}
}
