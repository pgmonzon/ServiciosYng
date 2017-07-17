package core

import (
  "time"
  "net/http"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/config"

  "github.com/dgrijalva/jwt-go"
  "github.com/dgrijalva/jwt-go/request"
)

func CrearToken(usuario models.Usuario) (string, error) {
  token := jwt.New(jwt.SigningMethodRS256)
  /*
  claims := make(jwt.MapClaims)
  claims["exp"] = time.Now().Add(time.Minute * config.ExpiraToken).Unix()
  claims["iat"] = time.Now().Unix()
  claims["usuarioID"] = usuario.ID
  token.Claims = claims
  */
  token.Claims = &models.TokenClaims{
    &jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Minute * config.ExpiraToken).Unix(),
    },
    usuario.ID,
  }

  tokenString, err := token.SignedString(config.SignKey)

  if err != nil {
    return "Error firmando el token", err
  }

  return tokenString, nil
}

func ValidarToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  start := time.Now()

	//token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
  token, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &models.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return config.VerifyKey, nil
		})

	if err == nil {
		if token.Valid {
      config.UsuarioActivoID = token.Claims.(*models.TokenClaims).UsuarioID
			next(w, r)
		} else {
      ErrorJSON(w, r, start, "Token inválido", http.StatusUnauthorized)
		}
	} else {
    ErrorJSON(w, r, start, "No tiene acceso a este recurso", http.StatusUnauthorized)
	}
}
