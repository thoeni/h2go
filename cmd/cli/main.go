package main

import (
	"flag"
	"fmt"
	"github.com/thoeni/h2go/h2go"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var w h2go.Waterer
	simulate := flag.Bool("simulate", false, "pass the -simulate flag to use an in memory implementation")
	set := flag.String("set", "", "pass the -set flag to use set a value [on/off]")
	flag.Parse()

	w = h2go.NewWaterer(*simulate)
	defer w.Close()

	// On SIGTERM stop pump and close
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		w.Stop()
		w.Close()
		os.Exit(1)
	}()

	switch *set {
	case "on":
		w.Start()
	case "off":
		w.Stop()
	default:
		fmt.Printf("Unknown set status: %s\n", *set)
	}
	fmt.Printf("Current waterer status: %s\n", boolToString(w.IsOn()))
}

func boolToString(b bool) string {
	s := "off"
	if b {
		s = "on"
	}
	return s
}