package main

import (
	"fmt"
	"log"
	"net/http"
	"web_app/src/router"
)

func main() {
	fmt.Println("Running Web_App")

	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))
}