package main

import (
	"gopooltask/gopooltask"
	"log"
	"sync"
	"time"
)

func main() {
	task := gopooltask.New(10)
	task.DebugLine = true

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		j := i
		wg.Add(1)
		task.AddTask(func() {
			//log.Println("BEGIN i=", j)
			time.Sleep(100 * time.Millisecond)
			log.Println("END   i=", j)
			wg.Done()
		})
	}
	wg.Wait()

}
