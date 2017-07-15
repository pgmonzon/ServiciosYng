package models

import (
  "time"

	"gopkg.in/mgo.v2/bson"
)

type Rol struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Rol       string        `json:"rol"`
  Estado    bool          `json:"estado"`
	Borrado		bool   				`json:"borrado"`
}

type RolID struct {
  ID        bson.ObjectId `json:"id"`
}

type Rol_adt struct {
  ID            bson.ObjectId `bson:"_id" json:"id"`
  RolID_adt     bson.ObjectId	`json:"rolID_adt"`
  Estado        bool          `json:"estado"`
	Borrado		    bool   				`json:"borrado"`
  Fecha_adt     time.Time     `json:"fecha_adt"`
  UsuarioID_adt bson.ObjectId	`json:"usuarioID_adt"`
  Oper_adt      string        `json: "oper_adt"`
}

type RolAlta struct {
	Rol       string        `json:"rol"`
  Estado    bool          `json:"estado"`
}
