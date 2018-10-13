package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/thoeni/h2go/h2go"
)

func main() {
	var w h2go.Waterer
	simulate := flag.Bool("simulate", false, "pass the -simulate flag to use an in memory implementation")
	flag.Parse()

	w = h2go.NewSystem(*simulate)
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

	http.Handle("/status", watererStatus(w))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type status struct {
	Status string `json:"status"`
}

func watererStatus(wt h2go.Waterer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getWatererStatus(w, r, wt)
		case http.MethodPost:
			postWatererStatus(w, r, wt)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func getWatererStatus(w http.ResponseWriter, r *http.Request, wt h2go.Waterer) {
	resp := status{boolToString(wt.IsOn())}
	b, _ := json.Marshal(resp)
	w.Write(b)
}

func postWatererStatus(w http.ResponseWriter, r *http.Request, wt h2go.Waterer) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	var req status
	json.Unmarshal(b, &req)
	log.Printf("Received request to update status to: %s\n", req.Status)
	switch req.Status {
	case "on":
		wt.Start()
		w.WriteHeader(http.StatusAccepted)
	case "off":
		wt.Stop()
		w.WriteHeader(http.StatusAccepted)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func boolToString(b bool) string {
	s := "off"
	if b {
		s = "on"
	}
	return s
}