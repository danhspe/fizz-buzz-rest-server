package health

var (
	MessageHealthy  = "healthy"
	MessageReady    = "ok"
	MessageNotReady = "not ok"
)

type Health interface {
	Healthy() (string, error)
	Ready() (string, error)
}
