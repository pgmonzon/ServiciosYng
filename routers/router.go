package routers

import (
	"net/http"
)

func InicializarRutas() {
	SetRutasUsuario()
	SetRutasRBAC()

	http.ListenAndServe(":3113", nil)
}
