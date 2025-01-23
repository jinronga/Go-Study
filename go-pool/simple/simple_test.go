package simple

import (
	"fmt"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	//创建一个Task
	task := NewTask(func() error {
		fmt.Println("创建一个Task:", time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := NewPool(3)

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for {
			p.EntryChannel <- task
		}
	}()

	//启动协程池p
	p.Start()
}
