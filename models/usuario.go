package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Usuario struct {
	ID        bson.ObjectId 	`bson:"_id" json:"id"`
	Usuario   string        	`json:"usuario"`
  Mail      string        	`json:"mail"`
	Clave			int64 					`json:"clave"`
	Estado    bool            `json:"estado"`
	Borrado   bool   				  `json:"borrado"`
	RolesID  	[]bson.ObjectId `json:"roles"`
}

type UsuarioID struct {
  ID        bson.ObjectId `json:"id"`
}

type Usuario_adt struct {
  ID            bson.ObjectId   `bson:"_id" json:"id"`
  Usuario_adt   bson.ObjectId	  `json:"usuario_adt"`
	Usuario			  string        	`json:"usuario"`
  Mail      		string        	`json:"mail"`
	Clave					int64 					`json:"clave"`
  Estado        bool            `json:"estado"`
	Borrado		    bool   				  `json:"borrado"`
  RolesID    		[]bson.ObjectId `json:"roles"`
  Fecha_adt     time.Time       `json:"fecha_adt"`
  UsuarioID_adt bson.ObjectId	  `json:"usuarioID_adt"`
  Oper_adt      string          `json: "oper_adt"`
}

type UsuarioAlta struct {
	Usuario			  string        	`json:"usuario"`
  Mail      		string        	`json:"mail"`
	Clave					string 					`json:"clave"`
  Estado        bool            `json:"estado"`
  RolesID    		[]bson.ObjectId `json:"roles"`
}

type UsuarioLogin struct {
	Usuario		string				`json:"usuario"`
	Clave			string				`json:"clave"`
}
