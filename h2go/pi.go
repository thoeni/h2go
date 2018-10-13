package h2go

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
)

type pi struct {
	WaterPump      rpio.Pin
	MoistureSensor rpio.Pin
}

func PiInit() (*pi, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}

	waterPin := rpio.Pin(18)
	waterPin.Output()

	moisturePin := rpio.Pin(12)
	moisturePin.Input()

	moisturePin.PullUp()
	moisturePin.Detect(rpio.AnyEdge)

	pi := pi{
		WaterPump: waterPin,
		MoistureSensor: moisturePin,
	}

	return &pi, nil
}

func (p *pi) IsOn() bool {
	if p.WaterPump.Read() == rpio.Low {
		return true
	}
	return false
}

func (p *pi) Stop() {
	p.WaterPump.High()
}

func (p *pi) Start() {
	p.WaterPump.Low()
}

func (p *pi) Toggle() {
	p.WaterPump.Toggle()
}

func (*pi) Close() error {
	fmt.Println("Closing memory mapping with RaspberryPI")
	return rpio.Close()
}

func (p *pi) MoistureSensorDetect() bool {
	return p.MoistureSensor.EdgeDetected()
}

func (p *pi) StopMoistureSensorDetection() {
	p.MoistureSensor.Detect(rpio.NoEdge)
}