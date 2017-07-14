package services

import (
  "log"
  "net/http"

  "github.com/pgmonzon/ServiciosYng/routers"
)

func ArrancarServer() {
  router := routers.InicializarRutas()
	log.Fatal(http.ListenAndServe(":3113", router))
}
