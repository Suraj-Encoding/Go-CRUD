package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Suraj-Encoding/api/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	fmt.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
