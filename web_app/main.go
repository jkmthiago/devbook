package main

import (
	"fmt"
	"log"
	"net/http"
	"web_app/src/config"
	"web_app/src/router"
	"web_app/src/utils"
)

func main() {
	fmt.Println("Charging prerequisites")
	config.LoadEnvVar()	

	fmt.Println("Charging Web App")
	utils.LoadTemplates()

	r := router.Generate()

	fmt.Printf("Web App is Running on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}