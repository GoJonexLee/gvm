package classfile

func readMembers(read *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := read.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(read, cp)
	}
	return members
}

func readMember(read *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     read.readUint16(),
		nameIndex:       read.readUint16(),
		descriptorIndex: read.readUint16(),
		attributes:      readAttributes(read, cp),
	}
}

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlags
}

func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}

func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (mi *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (mi *MemberInfo) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute)
		}
	}
	return nil
}

func (mi *MemberInfo) RuntimeVisibleAnnotationsAttributeData() []byte {
	return mi.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}

func (mi *MemberInfo) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return mi.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}

func (mi *MemberInfo) AnnotationDefaultAttributeData() []byte {
	return mi.getUnparsedAttributeData("AnnotationDefault")
}

func (mi *MemberInfo) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttrbute:
			unparsedAttr := attrInfo.(*unparsedAttr)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
