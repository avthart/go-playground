package main

import (
	//"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "192.168.99.100:8500"

	client, _ := api.NewClient(config)

	status := client.Status()
	leader, err := status.Leader()
	if err != nil {
		panic(err)
	}
	log.Println("consul: current leader ", leader)

	q := new(api.QueryOptions)
	q.WaitIndex = 0

	catalog := client.Catalog()
	services, _, err := catalog.Services(q)
	if err != nil {
		log.Println("error while retrieving services", err)
		return
	}
	for s := range services {
		log.Println("service: ", s)
	}
}
