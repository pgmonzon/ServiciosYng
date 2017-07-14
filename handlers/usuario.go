package handlers

import (
  "time"
  "encoding/json"
  "net/http"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/core"

  "gopkg.in/mgo.v2/bson"
)

// Valida las credenciales del usuario
func UsuarioLogin(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
  var usuarioLogin models.UsuarioLogin
  var usuario models.Usuario

  // Verifico que sea correcto el formato del JSON
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&usuarioLogin)
  if err != nil {
    core.ErrorJSON(w, r, start, "JSON decode erróneo", http.StatusBadRequest)
    return
  }

  // hago las validaciones de los campos obligatorios
	if usuarioLogin.Usuario == "" {
		core.ErrorJSON(w, r, start, "El usuario no puede estar vacío", http.StatusUnauthorized )
		return
	}
  if usuarioLogin.Clave == "" {
		core.ErrorJSON(w, r, start, "La clave no puede estar vacía", http.StatusUnauthorized)
		return
	}

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
	defer session.Close()

  // Intento traer el Usuario
	collection := session.DB("yangee").C("usuario")
	collection.Find(bson.M{"usuario": usuarioLogin.Usuario, "clave": core.HashSha512(usuarioLogin.Clave)}).One(&usuario)
	if usuario.ID == "" {
		core.ErrorJSON(w, r, start, "Acceso denegado", http.StatusUnauthorized)
	} else {
    token, err := core.CrearToken(usuario)
    if err != nil {
      core.ErrorJSON(w, r, start, token, http.StatusInternalServerError)
    }
    response, err := json.Marshal(models.Token{token})
		core.FatalErr(err)
		core.RespuestaJSON(w, r, start, response, http.StatusOK)
	}
}

// Lista todos los usuarios
func UsuarioListar(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
	var usuarios []models.Usuario

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
	defer session.Close()

  // Intento traer todos
	collection := session.DB("yangee").C("usuario")
	collection.Find(bson.M{}).All(&usuarios)
	response, err := json.Marshal(usuarios)
	core.FatalErr(err)
	core.RespuestaJSON(w, r, start, response, http.StatusOK)
}
