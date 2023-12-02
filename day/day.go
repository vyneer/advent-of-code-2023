package day

type Solver interface {
	Solve() ([]string, error)
	GetDayString() string
}
