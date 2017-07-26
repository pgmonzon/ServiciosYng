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

func validarNoExiste(rolAlta string) error {
  var rol models.Rol
  var errVal error
  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
  defer session.Close()

  // Me fijo si el rol ya existe
	collection := session.DB("yangee").C("rol")
	err := collection.Find(bson.M{"rol": bson.M{"$regex": rolAlta}}).One(&rol)
	if err != nil {
    if err.Error() == "not found" {
      return errVal
    } else {
      errVal = errors.New("Falló la validación del rol")
      return errVal
    }
  }
  errVal = errors.New("es duplicado")
  return errVal
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
    Assert: validarNoExiste(rolAlta.Rol),
    Insert: rol,
    }, {
      C: "rolAdt",
      Id: rolAdt.ID,
      Insert: rolAdt,
    }}

  err = runner.Run(ops, rol.ID, nil)
  if err != nil {
    if err.Error() == "Insert can only Assert txn.DocMissing%!(EXTRA *errors.errorString=es duplicado)" {
      core.ErrorJSON(w, r, start, "es duplicado", http.StatusBadRequest)
    } else {
      fmt.Println(err.Error())

      core.ErrorJSON(w, r, start, "error en la transacción", http.StatusBadRequest)
    }
    return
  }

  response, err := json.Marshal(models.RolID{rol.ID})
  core.FatalErr(err)
  core.RespuestaJSON(w, r, start, response, http.StatusCreated)
}
