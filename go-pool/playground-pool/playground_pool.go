package playground_pool

import (
	"fmt"
	"gopkg.in/go-playground/pool.v3"
	"time"
)

func SendMail(n int) pool.WorkFunc {
	fn := func(wu pool.WorkUnit) (interface{}, error) {

		// sleep 1s 模拟发邮件过程
		time.Sleep(time.Second * 1)
		// 模拟异常任务需要取消
		if n == 17 {
			wu.Cancel()
		}
		if wu.IsCancelled() {
			return false, nil
		}
		fmt.Println("send to", n)
		return true, nil
	}
	return fn
}
