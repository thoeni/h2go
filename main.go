package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thoeni/h2go/h2go"
)

func main() {
	var w h2go.Waterer
	simulate := flag.Bool("simulate", false, "pass the -simulate flag to use an in memory implementation")
	flag.Parse()

	w = newWaterer(*simulate)
	defer w.Close()

	fmt.Printf("Pump: %+v\n", w.IsOn())
	w.Toggle()
	fmt.Printf("Pump: %+v\n", w.IsOn())
}

func newWaterer(simulate bool) h2go.Waterer {
	switch simulate {
	case true:
		return h2go.InitSimulator()
	default:
		w, err := h2go.PiInit()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return w
	}
}
