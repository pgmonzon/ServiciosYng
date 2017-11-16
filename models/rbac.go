package models

import (
  "time"

	"gopkg.in/mgo.v2/bson"
)

type Rol struct {
	ID          bson.ObjectId   `bson:"_id" json:"id"`
	Rol         string          `json:"rol"`
  Estado      bool            `json:"estado"`
	Borrado		  bool   				  `json:"borrado"`
  RecursosID  []bson.ObjectId `json:"recursos"`
}

type RolID struct {
  ID        bson.ObjectId `json:"id"`
}

type Rol_adt struct {
  ID            bson.ObjectId   `bson:"_id" json:"id"`
  RolID_adt     bson.ObjectId	  `json:"rolID_adt"`
  Rol           string          `json:"rol"`
  Estado        bool            `json:"estado"`
	Borrado		    bool   				  `json:"borrado"`
  RecursosID    []bson.ObjectId `json:"recursos"`
  Fecha_adt     time.Time       `json:"fecha_adt"`
  UsuarioID_adt bson.ObjectId	  `json:"usuarioID_adt"`
  Oper_adt      string          `json: "oper_adt"`
}

type RolAlta struct {
	Rol         string          `json:"rol"`
  Estado      bool            `json:"estado"`
  RecursosID  []bson.ObjectId `json:"recursos"`
}

type Recurso struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Recurso   string        `json:"rol"`
  Estado    bool          `json:"estado"`
	Borrado		bool   				`json:"borrado"`
}

type RecursoID struct {
  ID        bson.ObjectId `json:"id"`
}

type Recurso_adt struct {
  ID            bson.ObjectId `bson:"_id" json:"id"`
  RecursoID_adt bson.ObjectId	`json:"recursoID_adt"`
  Recurso       string        `json:"recurso"`
  Estado        bool          `json:"estado"`
	Borrado		    bool   				`json:"borrado"`
  Fecha_adt     time.Time     `json:"fecha_adt"`
  UsuarioID_adt bson.ObjectId	`json:"usuarioID_adt"`
  Oper_adt      string        `json: "oper_adt"`
}

type RecursoAlta struct {
	Recurso   string        `json:"recurso"`
  Estado    bool          `json:"estado"`
}
