package go_ants

import (
	"github.com/panjf2000/ants"
	"sync"
	"testing"
)

func TestAnts(t *testing.T) {
	wg := sync.WaitGroup{}

	//申请一个协程池对象
	pool, _ := ants.NewPool(2)

	//关闭协程池
	defer pool.Release()

	// 向pool提交任务
	for i := 1; i <= 5; i++ {
		pool.Submit(sendMail(i, &wg))
		wg.Add(1)
	}
	wg.Wait()

}
