package core

import (
  "time"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/config"

  jwt "github.com/dgrijalva/jwt-go"
)

func CrearToken(usuario models.Usuario) (string, error) {
  token := jwt.New(jwt.SigningMethodRS256)
  claims := make(jwt.MapClaims)
  claims["exp"] = time.Now().Add(time.Minute * config.ExpiraToken).Unix()
  claims["iat"] = time.Now().Unix()
  claims["sub"] = usuario.ID
  token.Claims = claims

  tokenString, err := token.SignedString(config.SignKey)

  if err != nil {
    return "Error firmando el token", err
  }

  return tokenString, nil
}
