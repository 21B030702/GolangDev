package Assist

type Assistant struct {
	position string
	salary   float64
	address  string
}

func (a *Assistant) GetPosition() string {
	return a.position
}

func (a *Assistant) SetPosition(position string) {
	a.position = position
}

func (a *Assistant) GetSalary() float64 {
	return a.salary
}

func (a *Assistant) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Assistant) GetAddress() string {
	return a.address
}

func (a *Assistant) SetAddress(address string) {
	a.address = address
}

func NewEmployee(position string, salary float64, address string) *Assistant {
	return &Assistant{
		position: position,
		salary:   salary,
		address:  address,
	}
}
