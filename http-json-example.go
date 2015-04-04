package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// e.g. read spring boot metrics
	res, err := http.Get("http://localhost:8080/metrics")
	if err != nil {
		panic(err)
	}
	metricsData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var metrics map[string]float64
	json.Unmarshal(metricsData, &metrics)
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range metrics {
		fmt.Printf("%s: %f\n", key, value)
	}
}
