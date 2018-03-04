package constant

import (
	"gvm/instructions/base"

	"jvm.go/jvmgo/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (np *NOP) Excute(frame *rtda.Frame) {}
