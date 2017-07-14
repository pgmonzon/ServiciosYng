package main

import (
	"fmt"

	"github.com/pgmonzon/ServiciosYng/services"
	"github.com/pgmonzon/ServiciosYng/config"
)

func main() {
	fmt.Println("Yangee REST API Services...")

	config.Inicializar()
	services.ArrancarServer()
}
