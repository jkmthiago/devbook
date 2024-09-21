package main

import (
	// "encoding/hex"
	"fmt"
	"log"
	"net/http"
	"web_app/src/config"
	"web_app/src/cookies"
	"web_app/src/router"
	"web_app/src/utils"

	// "github.com/gorilla/securecookie"
)

// func init()  {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)

// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(blockKey)
// }

func main() {
	fmt.Println("Charging prerequisites")
	config.LoadEnvVar()	
	cookies.Config()

	fmt.Println("Charging Web App")
	utils.LoadTemplates()

	r := router.Generate()

	fmt.Printf("Web App is Running on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}