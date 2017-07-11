package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pgmonzon/ServiciosYng/routers"
	"github.com/pgmonzon/ServiciosYng/config"
)

func main() {
	fmt.Println("Yangee REST API Services...")

	config.Inicializar()
	router := routers.InicializarRutas()

	log.Fatal(http.ListenAndServe(":3113", router))
}
