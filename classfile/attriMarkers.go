package classfile

type MarkerAttribute struct{}

func (ma *MarkerAttribute) readInfo(reader *ClassReader) {

}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}
