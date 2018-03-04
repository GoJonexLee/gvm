package heap

// 局部变量表
type Slot struct {
	Val int64	// 存放整数
	Ref *Object	// 存放引用
}

func NewRefSlot(ref *Object) Slot {
	return Slot{0, ref)}
}

var EmptySlot = Slot{0, nil}