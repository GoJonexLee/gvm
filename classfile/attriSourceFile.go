package classfile

type SourceFileAttribute struct {
	cp              *ConstantPool
	SourceFileIndex uint16
}

func (sf *SourceFileAttribute) readInfo(reader *ClassReader) {
	sf.SourceFileIndex = reader.readUint16()
}

func (sf *SourceFileAttribute) FileName() string {
	return sf.cp.getUtf8(sf.SourceFileIndex)
}
