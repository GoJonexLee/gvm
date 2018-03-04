package rtda

import "jvm.go/jvmgo/rtda/heap"

type Frame struct {
	lower        *Frame
	thread       *Thread
	method       *heap.Method
	localVars    *LocalVars    // 局部变量表指针
	operandStack *OperandStack // 操作数栈指针
	maxLocals    uint
	maxStack     uint
	nextPC       int
	onPopAction  func()
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return *Frame{
		thread:       thread,
		method:       method,
		maxLocals:    method.MaxLocals(),
		maxStack:     method.MaxStack(),
		localVars:    newLocalVars(method.MaxLocals),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (f *Frame) reset(method *heap.Method) {
	f.method = method
	f.nextPC = 0
	f.lower = nil
	f.onPopAction = nil

	if f.maxLocals > 0 {
		f.localVars.clear()
	}

	if f.maxStack > 0 {
		f.operandStack.clear()
	}
}

func (f *Frame) Thread() *Thread {
	return f.thread
}
func (f *Frame) Method() *heap.Method {
	return f.method
}
func (f *Frame) LocalVars() *LocalVars {
	return f.localVars
}
func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}
func (f *Frame) NextPC() int {
	return f.nextPC
}
func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}
func (f *Frame) SetOnPopAction(f func()) {
	f.onPopAction = f
}

func (f *Frame) RevertNextPC() {
	f.nextPC = f.thread.pc
}

func (f *Frame) ClassLoader() *heap.ClassLoader {
	return f.BootLoader()
}
func (f *Frame) ConstantPool() *heap.ConstantPool {
	return f.method.ConstantPool()
}
