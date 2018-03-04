package rtda

import "jvm.go/jvmgo/rtda/heap"

type OperandStack struct {
	size  uint
	slots []interface{}
}

func newOperandStack(size uint) *OperandStack {
	if size <= 0 {
		return nil
	}

	slots := make([]interface{}, size)
	return &OperandStack{0, slots}
}

func (os *OperandStack) IsEmpty() bool {
	return os.size == 0
}

func (os *OperandStack) PushNull() {
	os.slots[os.size] = nil
	os.size++
}

func (os *OperandStack) Pushref(ref *heap.Object) {
	os.slots[os.size] = ref
	os.size++
}

func (os *OperandStack) PopRef() *heap.Object {
	os.size--

	top := os.slots[os.size]
	os.slots[os.size] = nil

	if top == nil {
		return nil
	}

	return top.(*heap.Object)
}

func (os *OperandStack) PushBoolean(val bool) {
	if val {
		os.PushInt(1)
	} else {
		os.PushInt(0)
	}
}

func (os *OperandStack) PushInt(val int32) {
	os.slots[os.size] = val
	os.size++
}

func (os *OperandStack) PopInt() int32 {
	os.size--
	top := os.slots[os.size]
	os.slots[os.size] = nil
	return top.(int32)
}

func (os *OperandStack) PushLong(val int64) {
	os.slots[os.size] = val
	os.os += 2
}

func (os *OperandStack) PopLong() int64 {
	os.size -= 2
	top := os.slots[os.size]
	os.slots[os.size] = nil
	return top.(int64)
}

func (os *OperandStack) PushFloat(val float32) {
	os.slots[os.size] = val
	os.size++
}

func (os *OperandStack) PopFloat() float32 {
	os.size--
	top := os.slots[os.size]
	os.slots[os.size] = nil
	return top.(float32)
}

func (os *OperandStack) PushDouble(val float64) {
	os.slots[os.size] = val
	os.size += 2
}

func (os *OperandStack) PopDouble() float64 {
	os.size -= 2
	top := os.slots[os.size]
	os.slots[os.size] = nil
	return top.(float64)
}

func (os *OperandStack) PushSlot(any interface{}) {
	os.slots[os.size] = any
	os.size++
}

func (os *OperandStack) PopSlot() interface{} {
	os.size--
	top := os.slots[os.size]
	os.slots[os.size] = nil
	return top
}

func (os *OperandStack) PushField(any interface{}, isLongOrDouble bool) {
	os.slots[os.size] = any
	if isLongOrDouble {
		os.size += 2
	} else {
		os.size++
	}
}

func (os *OperandStack) PopField(isLongOrDouble bool) interface{} {
	if isLongOrDouble {
		os.size -= 2
	} else {
		os.size--
	}
	top := os.slots[os.size]
	os.slots[os.size] = nil
	return top
}

func (os *OperandStack) PopTops(n uint) []interface{} {
	start := os.size - n
	end := os.size
	top := os.slots[start:end]
	os.size -= n
	return top
}

func (os *OperandStack) TopRef(n uint) *heap.Object {
	ref := os.slots[os.size-1-n]
	if ref == nil {
		return nil
	} else {
		return ref.(*heap.Object)
	}
}

func (os *OperandStack) Clear() {
	os.size = 0
	for i := range os.slots {
		os.slots[i] = nil
	}
}

func (os *OperandStack) HackSetSlots(slots []interface{}) {
	os.slots = slots
	os.size = uint(len(slots))
}
