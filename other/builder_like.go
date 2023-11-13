package other

type Employee struct {
	ID   string
	Name string
	Age  int
}

func (e *Employee) WithID(id string) *Employee {
	e.ID = id
	return e
}

func (e *Employee) WithName(name string) *Employee {
	e.Name = name
	return e
}

func (e *Employee) WithAge(age int) *Employee {
	e.Age = age
	return e
}

func NewEmployee() *Employee {
	return &Employee{}
}
