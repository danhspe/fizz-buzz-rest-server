package health

var (
	MessageOK     = "ok"
	MessageFailed = "failed"
)

type Health interface {
	Healthy() (string, error)
	Ready() (string, error)
}
