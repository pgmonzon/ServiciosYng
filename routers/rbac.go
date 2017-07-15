package routers

import (
  "net/http"

  "github.com/pgmonzon/ServiciosYng/handlers"
  "github.com/pgmonzon/ServiciosYng/core"

  "github.com/codegangsta/negroni"
)

func SetRutasRBAC() {
  http.Handle("/roles", negroni.New(
    negroni.HandlerFunc(core.ValidarToken),
    negroni.Wrap(http.HandlerFunc(handlers.RolAgregar)),
    ))
}
