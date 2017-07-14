package main

import (
	"fmt"

	"github.com/pgmonzon/ServiciosYng/routers"
	"github.com/pgmonzon/ServiciosYng/config"
)

func main() {
	fmt.Println("Yangee REST API Services...")

	config.Inicializar()
	routers.InicializarRutas()
}
