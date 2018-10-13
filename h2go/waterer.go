package h2go

import (
	"log"
	"os"
)

type System interface {
	Waterer
	MoistureDetector
}

type Waterer interface {
	IsOn() bool
	Stop()
	Start()
	Toggle()
	Close() error
}

type MoistureDetector interface {
	MoistureSensorDetect() bool
	StopMoistureSensorDetection()
}

func NewSystem(simulate bool) System {
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