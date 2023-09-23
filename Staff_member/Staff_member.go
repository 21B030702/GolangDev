package Staff_member

type Staff struct {
	position string
	salary   float64
	address  string
}

func (s *Staff) GetPosition() string {
	return s.position
}

func (s *Staff) SetPosition(position string) {
	s.position = position
}

func (s *Staff) GetSalary() float64 {
	return s.salary
}

func (s *Staff) SetSalary(salary float64) {
	s.salary = salary
}

func (s *Staff) GetAddress() string {
	return s.address
}

func (s *Staff) SetAddress(address string) {
	s.address = address
}

func NewStaff(position string, salary float64, address string) *Staff {
	return &Staff{
		position: position,
		salary:   salary,
		address:  address,
	}
}
