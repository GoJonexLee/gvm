package base

import "jvm.go/jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) // 提取操作数
	Execute(frame *rtda.Frame)            // 执行指令逻辑
}

type NoOperandsInstruction struct {
}

func (no *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

// 跳转指令, offset表示偏移量
type BranchInstruction struct {
	offset int
}

func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.offset = int(reader.ReadInt16())
}

// 存储和加载累指令
type Index8instruction struct {
	Index uint
}

func (ib *Index8instruction) FetchOperands(reader *BytecodeReader) {
	ib.Index = uint(reader.ReadUint8())
}

// 常量池索引
type Index16Instruction struct {
	Index uint
}

func (ii *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	ii.Index = uint(reader.ReadUint16)
}
