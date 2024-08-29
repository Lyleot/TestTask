package main

import (
	r "TestTask/app/api"
	"TestTask/app/conf"
	"TestTask/app/server"
	"net/http"
	"sync"

	_ "TestTask/docs"
)

func main() {
	config := conf.New()

	server.ParseParams(config)

	router := r.SetRouter(config)

	log := config.LOG()
	hp := config.HTTP()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	port := ":" + hp.HTTP_PORT

	go func() {
		http.ListenAndServe(port, router)
		log.Printf("Start on port " + port)
		wg.Done()
	}()

	wg.Wait()
}
