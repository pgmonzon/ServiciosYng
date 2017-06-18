package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pgmonzon/ServiciosYng/routers"
)

func main() {
	fmt.Printf("Yangee REST API Service\n")

	router := routers.NewRouter()
	
	log.Fatal(http.ListenAndServe(":3113", router))
}
