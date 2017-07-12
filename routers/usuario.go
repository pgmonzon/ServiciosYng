package routers

import (
  "github.com/pgmonzon/ServiciosYng/handlers"

  "github.com/gorilla/mux"
)

func SetRutasUsuario(router *mux.Router) *mux.Router {
  router.HandleFunc("/api/usuarios/login", handlers.UsuarioLogin).Methods("POST")
  //router.HandleFunc("/api/usuarios/logout", handlers.UsuarioLogout).Methods("GET")
  //tokenString := req.Header.Get("Authorization")
	return router
}
