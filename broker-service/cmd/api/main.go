package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "4000"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := svr.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
