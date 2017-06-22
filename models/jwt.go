package models

import (
	"gopkg.in/mgo.v2/bson"
  "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	jwt.StandardClaims
}
