package Worker

type Worker struct {
	position string
	salary   float64
	address  string
}

func (w *Worker) GetPosition() string {
	return w.position
}

func (w *Worker) SetPosition(position string) {
	w.position = position
}

func (w *Worker) GetSalary() float64 {
	return w.salary
}

func (w *Worker) SetSalary(salary float64) {
	w.salary = salary
}

func (w *Worker) GetAddress() string {
	return w.address
}

func (w *Worker) SetAddress(address string) {
	w.address = address
}

func NewWorker(position string, salary float64, address string) *Worker {
	return &Worker{
		position: position,
		salary:   salary,
		address:  address,
	}
}
