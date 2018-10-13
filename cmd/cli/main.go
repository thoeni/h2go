package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/thoeni/h2go/h2go"
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
	case "":
		// Return current value:
		fmt.Printf("Waterer on?: %v\n", w.IsOn())
	case "on":
		w.Start()
		fmt.Printf("Waterer on?: %v\n", w.IsOn())
	case "off":
		w.Stop()
		fmt.Printf("Waterer on?: %v\n", w.IsOn())
	default:
		fmt.Printf("Unknown set status: %s\n", *set)
	}
}
