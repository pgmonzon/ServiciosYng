package routers

import (
	"net/http"
)

func InicializarRutas() {
	SetRutasUsuario()

	http.ListenAndServe(":3113", nil)
}
