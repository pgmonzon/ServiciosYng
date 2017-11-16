package main

import (
	"fmt"

	"github.com/pgmonzon/ServiciosYng/routers"
	"github.com/pgmonzon/ServiciosYng/config"
)

func main() {
	fmt.Println("Yangee REST API Micro Services v1.0...")

	config.Inicializar()
	routers.InicializarRutas()
}
