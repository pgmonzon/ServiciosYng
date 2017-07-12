package core

import (
  "time"
  "io/ioutil"
  "fmt"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/config"

  jwt "github.com/dgrijalva/jwt-go"
)

func CrearToken(usuario models.Usuario) (interface{}) {
  // Creación del token
  token := jwt.New(jwt.SigningMethodHS256)

  // Asignación de Claims
  claims := make(jwt.MapClaims)
  claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.ExpiraToken)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = usuario.ID
  token.Claims = claims

  // Firmo el token
  signBytes, err := ioutil.ReadFile(config.PrivKeyPath)
  if err != nil {
    fmt.Println("Firma token paso 1")
  }
  FatalErr(err)

  signKey, err := jwt.ParseRSAPublicKeyFromPEM(signBytes)
  if err != nil {
    fmt.Println("Firma token paso 2")
  }
  FatalErr(err)

  tokenString, err := token.SignedString(signKey)
  if err != nil {
    fmt.Println("Firma token paso 3")
  }
  FatalErr(err)

  // Devuelvo el token
  jsonToken := map[string]string{"token": tokenString}
  return jsonToken
}
