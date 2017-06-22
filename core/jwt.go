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
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  // Firmo el token con secret
  signedToken, _ := token.SignedString([]byte(config.Secret))

  // Pongo el token en una cookie en el cliente
  //cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: config.ExpiraCookie, HttpOnly: true}
  //http.SetCookie(res, &cookie)

  // Devuelvo el token
  jsonToken := map[string]string{"token": signedToken}
  return jsonToken
}
