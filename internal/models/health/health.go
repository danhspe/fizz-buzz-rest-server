package health

type Health interface {
	Healthy() (string, error)
	Ready() (string, error)
}
