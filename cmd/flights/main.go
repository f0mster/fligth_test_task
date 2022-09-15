package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/f0mster/flights/gen/golang/proto"

	"github.com/f0mster/flights/internal/fligths/layers/domain"
	"github.com/f0mster/flights/internal/fligths/layers/interfaces/requsets"
)

func main() {
	host := parseFlags()
	startServer(host)
}

func parseFlags() string {
	host := flag.String("host", "listen on host", "")
	flag.Parse()
	if *host == "" {
		fmt.Println("host flag must be set")
		os.Exit(1)
	}
	return *host
}

func startServer(host string) {
	flt := domain.Flights{}
	fltController := requsets.NewFlightsController(&flt)
	flightService := proto.NewFlightsFinderServer(fltController)
	srv := &http.Server{
		Addr:    host,
		Handler: flightService,
	}
	fmt.Println("starting server on " + host)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("server stopped. error:", err)
	}
}
