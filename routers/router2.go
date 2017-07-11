package routers

import (
  "log"
	"net/http"

  "github.com/pgmonzon/ServiciosYng/handlers"

	"github.com/gorilla/mux"
)

// responde a las rutas no definidas
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%d",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		http.StatusNotFound,
		0,
		0,
	)
	w.WriteHeader(http.StatusNotFound)
}

func NewRouter() *mux.Router {
  r := mux.NewRouter().StrictSlash(false)
  //ejemplos
  /*
	r.HandleFunc("/api/todos", TodoIndex).Methods("GET")
	r.HandleFunc("/api/todos/{todoID}", TodoShow).Methods("GET")
	r.HandleFunc("/api/todos", TodoAdd).Methods("POST")
	r.HandleFunc("/api/todos/{todoID}", TodoUpdate).Methods("PUT")
	r.HandleFunc("/api/todos/{todoID}", TodoDelete).Methods("DELETE")
	r.HandleFunc("/api/todos/search/byname/{todoName}", TodoSearchName).Methods("GET")
	r.HandleFunc("/api/todos/search/bystatus/{status}", TodoSearchStatus).Methods("GET")
  */
  //handlers.usuario
  r.HandleFunc("/api/usuarios", handlers.UsuarioListar).Methods("GET")
  r.HandleFunc("/api/usuarios/{usuarioID}", handlers.UsuarioTraer).Methods("GET")
  r.HandleFunc("/api/usuarios", handlers.UsuarioRegistrar).Methods("POST")
  r.HandleFunc("/api/usuarios/{usuarioID}", handlers.UsuarioModificar).Methods("PUT")
  r.HandleFunc("/api/usuarios/{usuarioID}", handlers.UsuarioBorrar).Methods("DELETE")
  r.HandleFunc("/api/usuarios/buscar/porusuario/{usuarioUsuario}", handlers.UsuarioBuscarUsuario).Methods("GET")
  r.HandleFunc("/api/usuarios/login", handlers.UsuarioLogin).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(NotFound)
	return r
}
