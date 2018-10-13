package h2go

type Waterer interface {
	IsOn() bool
	Stop()
	Start()
	Toggle()
	Close() error
}
