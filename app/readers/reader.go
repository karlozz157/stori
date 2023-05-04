package readers

type Reader interface {
	GetData() ([][]string, error)
}
