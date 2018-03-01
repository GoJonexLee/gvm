package classfile

import (
	"fmt"
)

type ConstantPool struct {
	cfs []ConstantInfo
	cf  *ClassFile
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if id := cp[index]; id != nil {
		return id
	}
	panic(fmt.Errorf("Invalid constant pool index: %v!", id))
}

func (cp *ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ntInfo.nameIndex)
	tp := cp.getUtf8(ntInfo.descriptorIndex)
	return name, tp
}

func (cp *ConstantPool) getClassName(index uint16) string {
	name := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(name.nameIndex)
}

func (cp *ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func readConstantPool(read *ClassReader) *ConstantPool {
	cpCount := int(read.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(read, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return &ConstantPool{cfs: cp}
}
