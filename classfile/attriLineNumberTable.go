package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (lt *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberLength := reader.readUint16()
	lt.lineNumberTable = make([]*LineNumberTableEntry, lineNumberLength)

	for i := range lineNumberLength {
		lt.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    read.readUint16(),
			lineNumber: read.readUint16(),
		}
	}
}

func (lt *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(lt.lineNumberTable) - 1; i >= 0; i-- {
		entry := lt.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
