package models

import (
  "github.com/dgrijalva/jwt-go"
  "gopkg.in/mgo.v2/bson"
)

type Token struct {
  Token string `json:"token"`
}

type TokenClaims struct {
  *jwt.StandardClaims
  UsuarioID bson.ObjectId `json:"id"`
}
