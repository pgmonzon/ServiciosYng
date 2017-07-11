package core

import (
  "fmt"
  "time"

  "github.com/pgmonzon/ServiciosYng/models"
  "github.com/pgmonzon/ServiciosYng/config"

  jwt "github.com/dgrijalva/jwt-go"
)

func CrearToken(usuario models.Usuario) (interface{}) {
  // Asignación de Claims
  expiraToken := time.Now().Add(config.ExpiraToken * 1).Unix()
  claims := models.Claims {
    usuario.ID,
    jwt.StandardClaims {
      ExpiresAt: expiraToken,
    },
  }

  // Creación del token
  token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

  // Firmo el token con Secret
  signedToken, _ := token.SignedString([]byte(config.Secret))

  // Devuelvo el token
  jsonToken := map[string]string{"token": signedToken}
  return jsonToken
}

func ValidarToken(tokenString string) (interface{}) {
  fmt.Println(tokenString)
  token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return config.Secret, nil
	})

  if err != nil {
    return "qwerty"
  }

	claims := token.Claims.(*models.Claims)
	fmt.Println(claims.ID)
  return claims.ID

/*
  token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Método de firma inesperado")
    } else {
      return config.Secret, nil
    }

    if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
      return claims.ID, nil
    } else {
      return "asd", nil
    }
  })

  if err != nil {
    return "qwe"
  }
  return token
*/
}
