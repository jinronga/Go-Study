package simple

import "fmt"

type (
	Task struct {
		fn func() error // function to execute
	}

	Pool struct {
		EntryChannel chan *Task //对外接收Task的入口
		workerNum    int        //协程池最大worker数量,限定Goroutine的个数
		JobsChannel  chan *Task //协程池内部的任务就绪队列
	}
)

// NewTask creates a new Task with the given function to execute.
func NewTask(fn func() error) *Task {
	return &Task{fn: fn}
}

// Execute executes the task's function and returns any error it may have produced.
func (t *Task) Execute() error {
	return t.fn()
}

func NewPool(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		workerNum:    cap,
		JobsChannel:  make(chan *Task),
	}
}

func (p *Pool) worker(workerID int) {
	for task := range p.JobsChannel {

		task.Execute()
		// 任务执行完成后，将任务重新放回任务队列，等待再次被调度
		fmt.Println("worker ID ", workerID, " 执行完毕任务")
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.workerNum; i++ {
		fmt.Println("开启固定数量的Worker:", i)
		go p.worker(i)
	}

	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}

	close(p.JobsChannel)
	fmt.Println("执行完毕需要关闭JobsChannel")

	close(p.EntryChannel)
	fmt.Println("执行完毕需要关闭EntryChannel")

}
