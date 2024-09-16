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

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d", config.Api_port), r))
	fmt.Printf("Api is Running on port %d\n", config.Api_port)
}
