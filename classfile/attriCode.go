package classfile

type CodeAttribute struct {
	cp             *ConstantPool
	MaxStack       uint16                 // 操作数栈的最大深度
	MaxLocals      uint16                 // 局部变量表的大小
	Code           []byte                 // 字节码
	ExceptionTable []*ExceptionTableEntry // 异常处理表
	Attributes     AttributeTable         // 属性表
}

func (ca *CodeAttribute) readInfo(reader *ClassReader) {
	ca.MaxStack = reader.readUint16()
	ca.MaxLocals = reader.readUint16()
	ca.Code = reader.readBytes(reader.readUint32())
	ca.ExceptionTable = readExceptionTable(reader)
	ca.Attributes = readAttribute(reader, ca.cp)
}

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			StartPc:   reader.readUint16(),
			EndPc:     reader.readUint16(),
			HandlerPc: reader.readUint16(),
			CatchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
