package handlers

import (
  "time"
  "encoding/json"
  "net/http"
  "fmt"
  "errors"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/core"
  "github.com/pgmonzon/ServiciosYng/config"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "gopkg.in/mgo.v2/txn"
)

func validar() error {
  err := errors.New("es duplicado")
  return err
}

func RolAgregar(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var rolAlta models.RolAlta
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&rolAlta)
  if err != nil {
    core.ErrorJSON(w, r, start, "JSON decode erróneo", http.StatusBadRequest)
    return
  }

  // hago las validaciones de los campos obligatorios
	if rolAlta.Rol == "" {
		core.ErrorJSON(w, r, start, "El rol no puede estar vacío", http.StatusBadRequest)
		return
	}

  // establezco los campos que voy a guardar
  var rol models.Rol
  var rolAdt models.Rol_adt
	objID := bson.NewObjectId()
	rol.ID = objID
  rol.Rol = rolAlta.Rol
  rol.Estado = rolAlta.Estado
  rol.Borrado = false
  // auditoría
  objID_adt := bson.NewObjectId()
  rolAdt.ID = objID_adt
  rolAdt.RolID_adt = rol.ID
  rolAdt.Estado = rol.Estado
  rolAdt.Borrado = rol.Borrado
  rolAdt.UsuarioID_adt = config.UsuarioActivoID
  rolAdt.Oper_adt = "RolAgregar"

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
  defer session.Close()

  // Me aseguro el índice
  col := session.DB("yangee").C("rol")
  index := mgo.Index{
    Key:        []string{"rol"},
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
    C: "rol",
    Id: rol.ID,
    Assert: validar(),
    Insert: rol,
    }, {
      C: "rolAdt",
      Id: rolAdt.ID,
      Insert: rolAdt,
    }}

  err = runner.Run(ops, rol.ID, nil)
  if err != nil {
    fmt.Println(err.Error())
    if err.Error() == "Insert can only Assert txn.DocMissing%!(EXTRA *errors.errorString=es duplicado)" {
      core.ErrorJSON(w, r, start, "es duplicado", http.StatusBadRequest)
    } else {
      core.ErrorJSON(w, r, start, "error en la transacción", http.StatusBadRequest)
    }
    return
  }
/*
  // Intento el alta
	err = col.Insert(rol)
	if err != nil {
    if mgo.IsDup(err) {
      core.ErrorJSON(w, r, start, "El rol ya existe", http.StatusBadRequest)
  		return
    }
		core.ErrorJSON(w, r, start, "No se registró el rol", http.StatusInternalServerError)
		return
	}

  err = colAdt.Insert(rolAdt)
	if err != nil {
		core.ErrorJSON(w, r, start, "No se registró la auditoría del rol", http.StatusInternalServerError)
		return
	}
*/
  response, err := json.Marshal(models.RolID{rol.ID})
  core.FatalErr(err)
  core.RespuestaJSON(w, r, start, response, http.StatusCreated)
}
