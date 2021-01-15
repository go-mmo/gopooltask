package safeRecover

import "log"

func Wrap(name string, cb func()) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("---------------CatchPanic[" + name + "] BEGIN----------------")
			log.Println(e)
			log.Println("---------------CatchPanic[" + name + "] END----------------")
		}
	}()
	cb()
}
