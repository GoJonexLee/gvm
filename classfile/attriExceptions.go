package classfile

type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (es *ExceptionAttribute) readInfo(reader *ClassReader) {
	es.exceptionIndexTable = reader.readUint16s()
}

func (es *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return es.exceptionIndexTable
}
