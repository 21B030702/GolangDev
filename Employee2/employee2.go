package Employee2

type Employee struct {
	position string
	salary   float64
	address  string
}

func (e *Employee) GetPosition() string {
	return e.position
}

func (e *Employee) SetPosition(position string) {
	e.position = position
}

func (e *Employee) GetSalary() float64 {
	return e.salary
}

func (e *Employee) SetSalary(salary float64) {
	e.salary = salary
}

func (e *Employee) GetAddress() string {
	return e.address
}

func (e *Employee) SetAddress(address string) {
	e.address = address
}

func NewEmployee(position string, salary float64, address string) *Employee {
	return &Employee{
		position: position,
		salary:   salary,
		address:  address,
	}
}
