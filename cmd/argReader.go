package cmd

type ArgReader struct {
	args []string
}

func (ar *ArgReader) hasMoreOptions() bool {
	return len(ar.args) > 0 && ar.args[0][0] == '-'
}

func (ar *ArgReader) removeFirse() string {
	first := ar.args[0]
	ar.args = ar.args[1:]
	return first
}
