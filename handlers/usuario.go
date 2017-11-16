package handlers

import (
  "time"
  "encoding/json"
  "net/http"
  "errors"
  "fmt"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/core"
  "github.com/pgmonzon/ServiciosYng/config"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "gopkg.in/mgo.v2/txn"
)

func noExisteUsuario(usuarioAlta string) error {
  var usuario models.Usuario
  var errVal error
  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
  defer session.Close()

  // Me fijo si el rol ya existe
	collection := session.DB("yangee").C("usuario")
	err := collection.Find(bson.M{"usuario": bson.M{"$regex": usuarioAlta}}).One(&usuario)
	if err != nil {
    if err.Error() == "not found" {
      return errVal
    } else {
      errVal = errors.New("Falló la validación del usuario")
      return errVal
    }
  }
  errVal = errors.New("es duplicado")
  return errVal
}

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

func UsuarioAgregar(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuarioAlta models.UsuarioAlta
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&usuarioAlta)
  if err != nil {
    core.ErrorJSON(w, r, start, "JSON decode erróneo", http.StatusBadRequest)
    return
  }

  // hago las validaciones de los campos obligatorios
	if usuarioAlta.Usuario == "" {
		core.ErrorJSON(w, r, start, "El usuario no puede estar vacío", http.StatusBadRequest)
		return
	}
  if usuarioAlta.Mail == "" {
		core.ErrorJSON(w, r, start, "El mail no puede estar vacío", http.StatusBadRequest)
		return
	}
  if usuarioAlta.Clave == "" {
		core.ErrorJSON(w, r, start, "La clave no puede estar vacía", http.StatusBadRequest)
		return
	}

  // establezco los campos que voy a guardar
  var usuario models.Usuario
  var usuarioAdt models.Usuario_adt
	objID := bson.NewObjectId()
	usuario.ID = objID
  usuario.Usuario = usuarioAlta.Usuario
  usuario.Mail = usuarioAlta.Mail
  usuario.Clave = core.HashSha512(usuarioAlta.Clave)
  usuario.Estado = usuarioAlta.Estado
  usuario.Borrado = false
  usuario.RolesID = usuarioAlta.RolesID
  // auditoría
  objID_adt := bson.NewObjectId()
  usuarioAdt.ID = objID_adt
  usuarioAdt.Usuario_adt = usuario.ID
  usuarioAdt.Usuario = usuario.Usuario
  usuarioAdt.Mail = usuario.Mail
  usuarioAdt.Clave = usuario.Clave
  usuarioAdt.Estado = usuario.Estado
  usuarioAdt.Borrado = usuario.Borrado
  usuarioAdt.RolesID = usuario.RolesID
  usuarioAdt.UsuarioID_adt = config.UsuarioActivoID
  usuarioAdt.Oper_adt = "UsuarioAgregar"

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
  defer session.Close()

  // Me aseguro el índice
  col := session.DB("yangee").C("usuario")
  index := mgo.Index{
    Key:        []string{"usuario"},
    Unique:     true,
    DropDups:   true,
    Background: true,
    Sparse:     true,
  }
  err = col.EnsureIndex(index)
  if err != nil {
    panic(err)
  }

  runner := txn.NewRunner(session.DB(config.DB_Name).C(config.DB_Transaction))
  ops := []txn.Op{{
    C: "usuario",
    Id: usuario.ID,
    Assert: noExisteUsuario(usuarioAlta.Usuario),
    Insert: usuario,
    }, {
      C: "usuarioAdt",
      Id: usuarioAdt.ID,
      Insert: usuarioAdt,
    }}

  err = runner.Run(ops, usuario.ID, nil)
  if err != nil {
    if err.Error() == "Insert can only Assert txn.DocMissing%!(EXTRA *errors.errorString=es duplicado)" {
      core.ErrorJSON(w, r, start, "es duplicado", http.StatusBadRequest)
    } else {
      fmt.Println(err.Error())
      fmt.Println(usuario.ID)
      fmt.Println(usuarioAdt.ID)
      fmt.Println(usuarioAdt.Usuario_adt)
      fmt.Println(usuario.RolesID)
      fmt.Println(usuarioAdt.RolesID)
      core.ErrorJSON(w, r, start, "error en la transacción", http.StatusBadRequest)
    }
    return
  }

  response, err := json.Marshal(models.UsuarioID{usuario.ID})
  core.FatalErr(err)
  core.RespuestaJSON(w, r, start, response, http.StatusCreated)
}
