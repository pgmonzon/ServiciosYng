package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Usuario struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Usuario   string        `json:"usuario"`
  Mail      string        `json:"mail"`
	Clave			int64 				`json:"clave"`
}

type UsuarioLogin struct {
	Usuario		string				`json:"usuario"`
	Clave			string				`json:"clave"`
}
