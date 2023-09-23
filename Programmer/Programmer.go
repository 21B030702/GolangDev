package Programmer

type Programmer struct {
	position string
	salary   float64
	address  string
}

func (p *Programmer) GetPosition() string {
	return p.position
}

func (p *Programmer) SetPosition(position string) {
	p.position = position
}

func (p *Programmer) GetSalary() float64 {
	return p.salary
}

func (p *Programmer) SetSalary(salary float64) {
	p.salary = salary
}

func (p *Programmer) GetAddress() string {
	return p.address
}

func (p *Programmer) SetAddress(address string) {
	p.address = address
}

func NewProgrammer(position string, salary float64, address string) *Programmer {
	return &Programmer{
		position: position,
		salary:   salary,
		address:  address,
	}
}
