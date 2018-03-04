package rtda

import (
	"sync"

	"jvm.go/jvmgo/options"
	"jvm.go/jvmgo/rtda/heap"
)

type Thread struct {
	pc              int
	stack           *Stack
	frameCache      *FrameCache
	jThread         *heap.Object
	lock            *sync.Mutex
	ch              chan int
	sleepingFlag    bool
	interruptedFlag bool
	parkingFlag     bool
	unparkedFlag    bool
}

func NewThread(jThread *heap.Object) *Thread {
	stack := newStack(options.ThreadStackSize)
	thread := &Thread{
		stack: stack,
		jThread: jThread,
		lock: &sync.Mutex{},
		ch: make(chan int)
	}
	thread.frameCache = newFrameCache(thread, 16)
	return thread
}
