package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Charging prerequisites")
	config.LoadEnvVar()

	fmt.Println("Charging Api")

	r := router.GenerateNewRoute()

	fmt.Printf("Api is Running on port %s\n", config.Api_port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Api_port), r))
}
