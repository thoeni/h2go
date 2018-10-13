package h2go

import (
	"log"
	"os"
)

type Waterer interface {
	IsOn() bool
	Stop()
	Start()
	Toggle()
	Close() error
}

func NewWaterer(simulate bool) Waterer {
	switch simulate {
	case true:
		return InitSimulator()
	default:
		w, err := PiInit()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		return w
	}
}