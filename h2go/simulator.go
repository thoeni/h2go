package h2go

type simulator struct{
	waterPumpOff bool
}

func InitSimulator() *simulator {
	return &simulator{}
}

func (s *simulator) IsOn() bool {
	return !s.waterPumpOff
}

func (s *simulator) Stop() {
	s.waterPumpOff = true
}

func (s *simulator) Start() {
	s.waterPumpOff = false
}

func (s *simulator) Toggle() {
	s.waterPumpOff = !s.waterPumpOff
}

func (simulator) Close() error {
	return nil
}

