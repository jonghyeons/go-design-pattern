package iterator

type Iterator interface {
	HasNext() bool
	Next() MenuItem
}

type MenuItem struct{}
