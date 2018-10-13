package h2go

import (
	"fmt"
	"time"
)

type simulator struct{
	waterPumpOff bool
}

func (s *simulator) MoistureSensorDetect() bool {
	panic("implement me")
}

func (s *simulator) StopMoistureSensorDetection() {
	panic("implement me")
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
	time.Sleep(5 * time.Second)
	s.waterPumpOff = !s.waterPumpOff
}

func (simulator) Close() error {
	fmt.Println("Closing water simulator")
	return nil
}

