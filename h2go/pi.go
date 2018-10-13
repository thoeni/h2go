package h2go

import "github.com/stianeikeland/go-rpio"

type pi struct {
	WaterPump rpio.Pin
}

func PiInit() (*pi, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}

	waterPin := rpio.Pin(18)
	waterPin.Output()

	pi := pi{
		WaterPump: waterPin,
	}

	return &pi, nil
}

func (p *pi) IsOn() bool {
	if p.WaterPump.Read() == rpio.Low {
		return false
	}
	return true
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
	return rpio.Close()
}
