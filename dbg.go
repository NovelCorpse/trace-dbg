package dbg

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

func Trace2(msg string) {
	pc, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)
	f := runtime.FuncForPC(pc)

	fmt.Printf("[%s][%s][%s]:[%d] [%s]\n", time.Now().Format("2006.01.02"), fileName, f.Name(), line, msg)
}

func Trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	_, fileName := path.Split(file)
	fmt.Printf("(full)[%s] [%s] : [%d] [%s]\n", file, fileName, line, f.Name())
}
