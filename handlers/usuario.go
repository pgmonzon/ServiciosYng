package handlers

import (
  "time"
  "encoding/json"
  "net/http"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/core"

  "gopkg.in/mgo.v2/bson"
  "github.com/gorilla/mux"
)

var null = "null"

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
	response, err := json.MarshalIndent(usuarios, "", "    ")
	if err != nil {
    panic(err)
	}
	core.RespuestaJSON(w, r, start, response, http.StatusOK)
}

// Devuelve un usuario por ID
func UsuarioTraer(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  var usuario models.Usuario

  // Verifico el formato del campo ID
  vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["usuarioID"]) != true {
		core.ErrorJSON(w, r, start, "Formato de ID incorrecto", http.StatusBadRequest)
		return
	}
	usuarioID := bson.ObjectIdHex(vars["usuarioID"])

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
	defer session.Close()

  // Intento traer el ID
	collection := session.DB("yangee").C("usuario")
	collection.Find(bson.M{"_id": usuarioID}).One(&usuario)
	if usuario.ID == "" {
		core.ErrorJSON(w, r, start, "Usuario no encontrado", http.StatusNotFound)
	} else {
		response, err := json.MarshalIndent(usuario, "", "    ")
		if err != nil {
			panic(err)
		}
		core.RespuestaJSON(w, r, start, response, http.StatusOK)
	}
}

// Registra un nuevo Usuario
func UsuarioRegistrar(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuarioRegistro models.UsuarioRegisro
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&usuarioRegistro)
  if err != nil {
    core.ErrorJSON(w, r, start, "JSON decode erróneo", http.StatusBadRequest)
    return
  }

  // hago las validaciones de los campos obligatorios
	if usuarioRegistro.Usuario == "" {
		core.ErrorJSON(w, r, start, "El usuario no puede estar vacío", http.StatusBadRequest)
		return
	}

  // establezco los campos que voy a guardar
  var usuario models.Usuario
	objID := bson.NewObjectId()
	usuario.ID = objID
  usuario.Usuario = usuarioRegistro.Usuario
  usuario.Mail = usuarioRegistro.Mail
  usuario.Clave = core.HashSha512(usuarioRegistro.Clave)

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
  defer session.Close()

  // Intento el alta
  collection := session.DB("yangee").C("usuario")
	err = collection.Insert(usuario)
	if err != nil {
		core.ErrorJSON(w, r, start, "No se registró el usuario", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(usuario.ID.Hex()))
	core.RespuestaJSON(w, r, start, []byte{}, http.StatusCreated)
}

// Modifica un usuario por ID
func UsuarioModificar(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
  var usuario models.Usuario

  // Verifico el formato del campo ID
	vars := mux.Vars(r)
  if bson.IsObjectIdHex(vars["usuarioID"]) != true {
		core.ErrorJSON(w, r, start, "Formato de ID incorrecto", http.StatusBadRequest)
		return
	}

  // Verifico el formato del JSON
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&usuario)
  if err != nil {
    core.ErrorJSON(w, r, start, "JSON decode erróneo", http.StatusBadRequest)
    return
  }

	usuarioID := bson.ObjectIdHex(vars["usuarioID"])

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
	defer session.Close()

  // Intento modificar
	collection := session.DB("yangee").C("usuario")
	err = collection.Update(bson.M{"_id": usuarioID},
		bson.M{"$set": bson.M{"usuario": usuario.Usuario, "mail": usuario.Mail}})
	if err != nil {
		core.ErrorJSON(w, r, start, "No se pudo encontrar el usuario "+string(usuarioID.Hex())+" para modificar", http.StatusNotFound)
		return
	}
	core.RespuestaJSON(w, r, start, []byte{}, http.StatusNoContent)
}

// Borra un usuario por ID
func UsuarioBorrar(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

  // Verifico el formato del campo ID
  vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["usuarioID"]) != true {
		core.ErrorJSON(w, r, start, "Formato de ID incorrecto", http.StatusBadRequest)
		return
	}
	usuarioID := bson.ObjectIdHex(vars["usuarioID"])

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
	defer session.Close()

  // Intento borrar
	collection := session.DB("yangee").C("usuario")
	err := collection.Remove(bson.M{"_id": usuarioID})
	if err != nil {
		core.ErrorJSON(w, r, start, "No se pudo encontrar el usuario "+string(usuarioID.Hex())+" para borrar", http.StatusNotFound)
		return
	}
	core.RespuestaJSON(w, r, start, []byte{}, http.StatusNoContent)
}

// Devuelve un usuario por usuario
func UsuarioBuscarUsuario(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuario []models.Usuario

	vars := mux.Vars(r)
	usuarioUsuario := vars["usuarioUsuario"]

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
	defer session.Close()

  // Intento traer el usuario
	collection := session.DB("yangee").C("usuario")
	err := collection.Find(bson.M{"usuario": bson.M{"$regex": usuarioUsuario}}).All(&usuario)
	if err != nil {
		core.ErrorJSON(w, r, start, "Falló la búsqueda por usuario", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(usuario, "", "    ")
	if err != nil {
		panic(err)
	}
	if string(response) == null {
		core.ErrorJSON(w, r, start, "No se encontró ningún usuario conteniendo "+usuarioUsuario, http.StatusNotFound)
		return
	}
	core.RespuestaJSON(w, r, start, response, http.StatusOK)
}
