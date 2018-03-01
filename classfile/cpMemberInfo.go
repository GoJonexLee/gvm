package classfile

type ConstantFielderInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMemberrefInfo struct {
	cp               *ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (cm *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	cm.classIndex = reader.readUint16()
	cm.nameAndTypeIndex = reader.readUint16()
}

func (cm *ConstantMemberrefInfo) ClassName() string {
	return cm.cp.getClassName(cm.classIndex)
}

func (cm *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return cm.cp.getNameAndType(cm.nameAndTypeIndex)
}
