package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (cn *ConstantNameAndTypeInfo) readInfo(read *ClassReader) {
	cn.nameIndex = read.readUint16()
	cn.descriptorIndex = read.readUint16
}
