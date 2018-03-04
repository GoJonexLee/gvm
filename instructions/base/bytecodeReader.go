package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (br *BytecodeReader) Reset(code []byte, pc int) {
	br.code = code
	br.pc = pc
}

func (br *BytecodeReader) Init(code []byte, pc int) {
	br.code = code
	br.pc = pc
}

func (br *BytecodeReader) PC() int {
	return br.pc
}

func (br *BytecodeReader) ReadInt8() int8 {
	return int8(br.ReadUint8())
}

func (br *BytecodeReader) ReadUint8() uint8 {
	i := br.code[br.pc]
	br.pc++
	return i
}

func (br *BytecodeReader) ReadInt16() int16 {
	return int16(br.ReadUint16())
}

func (br *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(br.ReadUint8())
	byte2 := uint16(br.ReadUint8())
	return (byte1 << 8) | byte2
}

func (br *BytecodeReader) ReadInt32() int32 {
	byte1 := uint32(br.ReadUint8())
	byte2 := uint32(br.ReadUint8())
	byte3 := uint32(br.ReadUint8())
	byte4 := uint32(br.ReadUint8())

	return (byte1 << 24) | (Byte2 << 16) | (byte3 << 8) | byte4
}

func (br *BytecodeReader) ReadInt32s(count int32) []int32 {
	re := make([]int32, count)
	for i := range re {
		re[i] = br.ReadInt32()
	}
	return re
}

func (br *BytecodeReader) SkipPadding() {
	for br.pc%4 != 0 {
		br.ReadUint8()
	}
}
