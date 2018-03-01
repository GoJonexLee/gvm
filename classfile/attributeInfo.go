package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp *ConstantPool) []AttributeInfo {
	attributes := make([]AttributeInfo, reader.readUint16())

	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes
}

func readAttribute(reader *ClassReader, cp *ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrLen := reader.readUint32()
	attrName := cp.getUtf8(attrNameIndex)
	attrInfo := newAttributeInfo(attrName, cp)

	if attrInfo == nil {
		attrInfo = &UnparsedAttribute{
			name:   attrName,
			length: attrLen,
		}
	}

	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, cp *ConstantPool) AttributeInfo {
	switch attrName {
	case "BootstrapMethod":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return _attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp: cp}
	case "Exception":
		return &ExceptionsAttribute{cp: cp}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "SourceFile":
		return &SourceTureAttribute{cp: cp}
	case "Synthetic":
		return _attrSynthetic
	default:
		return nil
	}
}
