package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jasonkeene/expmetric"
)

var (
	counter = expmetric.NewCounter("some_counter")
	gauge   = expmetric.NewGauge("some_gauge", "kWh")
)

func main() {
	go func() {
		for {
			counter.Add(1)
			gauge.Set(17234.235)
			time.Sleep(time.Second)
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
