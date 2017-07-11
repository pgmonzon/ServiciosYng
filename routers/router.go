package routers

import (
	"github.com/gorilla/mux"
)

func InicializarRutas() *mux.Router {
	router := mux.NewRouter()
	router = SetRutasUsuario(router)
	return router
}
