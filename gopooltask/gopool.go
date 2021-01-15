package gopooltask

import (
	"fmt"
	"gopooltask/safeRecover"
	"log"
	"sync/atomic"
)

type gopoolTask struct {
	maxTask   int
	totalLoop int32
	tasks     chan func()
	done      chan struct{}

	DebugLine bool //一轮任务结束,打印一条线
}

func New(maxTask int) *gopoolTask {
	this := &gopoolTask{maxTask: maxTask}
	this.tasks = make(chan func(), maxTask)
	this.done = make(chan struct{})
	for i := 0; i < maxTask; i++ {
		j := i
		go func() {
			for {
				select {
				case task := <-this.tasks:
					safeRecover.Wrap(fmt.Sprintf("gopoolTask[%v] Running", j), task)
					curLoop := atomic.AddInt32(&this.totalLoop, 1)
					if this.DebugLine && curLoop%10 == 0 && curLoop != 0 {
						log.Println("---------------------")
					}
				case <-this.done:
					return
				}
			}
		}()
	}
	return this
}
func (this *gopoolTask) AddTask(cb func()) {
	this.tasks <- cb
}

func (this *gopoolTask) Done() {
	close(this.done)
}
