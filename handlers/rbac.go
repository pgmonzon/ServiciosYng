package handlers

import (
  "time"
  "encoding/json"
  "net/http"
  "errors"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/core"
  "github.com/pgmonzon/ServiciosYng/config"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "gopkg.in/mgo.v2/txn"
)

func noExisteRol(rolAlta string) error {
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

func noExisteRecurso(recursoAlta string) error {
  var recurso models.Recurso
  var errVal error
  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
  defer session.Close()

  // Me fijo si el rol ya existe
	collection := session.DB("yangee").C("recurso")
	err := collection.Find(bson.M{"recurso": bson.M{"$regex": recursoAlta}}).One(&recurso)
	if err != nil {
    if err.Error() == "not found" {
      return errVal
    } else {
      errVal = errors.New("Falló la validación del recurso")
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
  rol.RecursosID = rolAlta.RecursosID
  // auditoría
  objID_adt := bson.NewObjectId()
  rolAdt.ID = objID_adt
  rolAdt.RolID_adt = rol.ID
  rolAdt.Rol = rol.Rol
  rolAdt.Estado = rol.Estado
  rolAdt.Borrado = rol.Borrado
  rolAdt.RecursosID = rol.RecursosID
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
    Assert: noExisteRol(rolAlta.Rol),
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
      core.ErrorJSON(w, r, start, "error en la transacción", http.StatusBadRequest)
    }
    return
  }

  response, err := json.Marshal(models.RolID{rol.ID})
  core.FatalErr(err)
  core.RespuestaJSON(w, r, start, response, http.StatusCreated)
}

func RecursoAgregar(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var recursoAlta models.RecursoAlta
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&recursoAlta)
  if err != nil {
    core.ErrorJSON(w, r, start, "JSON decode erróneo", http.StatusBadRequest)
    return
  }

  // hago las validaciones de los campos obligatorios
	if recursoAlta.Recurso == "" {
		core.ErrorJSON(w, r, start, "El recurso no puede estar vacío", http.StatusBadRequest)
		return
	}

  // establezco los campos que voy a guardar
  var recurso models.Recurso
  var recursoAdt models.Recurso_adt
	objID := bson.NewObjectId()
	recurso.ID = objID
  recurso.Recurso = recursoAlta.Recurso
  recurso.Estado = recursoAlta.Estado
  recurso.Borrado = false
  // auditoría
  objID_adt := bson.NewObjectId()
  recursoAdt.ID = objID_adt
  recursoAdt.RecursoID_adt = recurso.ID
  recursoAdt.Recurso = recurso.Recurso
  recursoAdt.Estado = recurso.Estado
  recursoAdt.Borrado = recurso.Borrado
  recursoAdt.UsuarioID_adt = config.UsuarioActivoID
  recursoAdt.Oper_adt = "RecursoAgregar"

  // Genero una nueva sesión Mongo
	session := core.GetMongoSession()
  defer session.Close()

  // Me aseguro el índice
  col := session.DB("yangee").C("recurso")
  index := mgo.Index{
    Key:        []string{"recurso"},
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
    C: "recurso",
    Id: recurso.ID,
    Assert: noExisteRecurso(recursoAlta.Recurso),
    Insert: recurso,
    }, {
      C: "recursoAdt",
      Id: recursoAdt.ID,
      Insert: recursoAdt,
    }}

  err = runner.Run(ops, recurso.ID, nil)
  if err != nil {
    if err.Error() == "Insert can only Assert txn.DocMissing%!(EXTRA *errors.errorString=es duplicado)" {
      core.ErrorJSON(w, r, start, "es duplicado", http.StatusBadRequest)
    } else {
      core.ErrorJSON(w, r, start, "error en la transacción", http.StatusBadRequest)
    }
    return
  }

  response, err := json.Marshal(models.RecursoID{recurso.ID})
  core.FatalErr(err)
  core.RespuestaJSON(w, r, start, response, http.StatusCreated)
}
