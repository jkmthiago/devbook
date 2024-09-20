package main

import (
	"fmt"
	"log"
	"net/http"
	"web_app/src/router"
	"web_app/src/utils"
)

func main() {
	// fmt.Println("Charging prerequisites")
	// config.LoadEnvVar()	

	fmt.Println("Charging Web App")
	utils.LoadTemplates()

	r := router.Generate()

	fmt.Println("Web App is Running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}