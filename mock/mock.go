package mock

type entry struct {
	vals []interface{}
	err error
}

type methodEntry struct {
	args []interface{}
}

type Mock struct {
	methods[string]*entry
	currentMethod string
}
