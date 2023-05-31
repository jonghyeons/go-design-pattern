package singleton

type Something struct {
	// ...
}

var uniqueInstance *Something

type Singleton struct{}

func (s *Singleton) GetInstance() *Something {
	if uniqueInstance == nil {
		return new(Something)
	}
	return uniqueInstance
}
