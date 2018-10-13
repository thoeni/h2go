package main

import (
	"flag"
	"fmt"
	"github.com/thoeni/h2go/h2go"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var s h2go.System
	simulate := flag.Bool("simulate", false, "pass the -simulate flag to use an in memory implementation")
	set := flag.String("set", "", "pass the -set flag to use set a value [on/off]")
	moisture := flag.Bool("moisture", false, "pass the -moisture flag to use check moisture level")
	flag.Parse()

	s = h2go.NewSystem(*simulate)
	defer s.Close()

	// On SIGTERM stop pump and close
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		s.Stop()
		s.Close()
		os.Exit(1)
	}()

	switch *set {
	case "on":
		s.Start()
	case "off":
		s.Stop()
	default:
	}
	fmt.Printf("Current waterer status: %s\n", boolToString(s.IsOn()))

	if *moisture {
		for i := 0; i < 2; {
			if s.MoistureSensorDetect() { // check if event occured
				fmt.Println("value changed")
				i++
			}
			time.Sleep(time.Second)
		}
	}
}

func boolToString(b bool) string {
	s := "off"
	if b {
		s = "on"
	}
	return s
}