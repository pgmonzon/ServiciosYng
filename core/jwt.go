package core

import (
  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/config"

  "github.com/dgrijalva/jwt-go"
)

func CrearToken(usuario models.Usuario) (interface{}) {
  // Asignación de Claims
  claims := models.Claims {
    usuario.ID,
    jwt.StandardClaims {
      ExpiresAt: config.ExpiraToken,
    },
  }

  // Creación del token
  token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

  // Firmo el token con secret
  signedToken, _ := token.SignedString([]byte(config.Secret))

  // Devuelvo el token
  jsonToken := map[string]string{"token": signedToken}
  return jsonToken
}
