package routers

import (
  "net/http"

  "github.com/pgmonzon/ServiciosYng/handlers"
  "github.com/pgmonzon/ServiciosYng/core"

  "github.com/codegangsta/negroni"
)

func SetRutasUsuario() {
  http.HandleFunc("/login", handlers.UsuarioLogin)
  http.HandleFunc("/usuario", handlers.UsuarioAgregar)

  http.Handle("/usuarios", negroni.New(
    negroni.HandlerFunc(core.ValidarToken),
    negroni.Wrap(http.HandlerFunc(handlers.UsuarioListar)),
    ))
}
