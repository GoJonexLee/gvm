package base

import "jvm.go/jvmgo/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thead().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
