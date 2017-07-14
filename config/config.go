package config

import (
  "io/ioutil"
  "log"

  "crypto/rsa"
  "github.com/dgrijalva/jwt-go"
)

const(
    // Base de datos
    DB_Host = "localhost"
    //DB_Host = "mongodb://127.0.0.1:27017"
    //DB_Host = "mongodb://yng_user:laser@ds021326.mlab.com:21326/yangee"
    DB_Name = "yangee"
    DB_User = "yng_Usr"
    DB_Pass = "1962Laser"

    // jwt
    privKeyPath = "C:/Users/Patricio/Google Drive/proyectoYangee/codigoGo/src/github.com/pgmonzon/ServiciosYng/config/keys/app.rsa"
    pubKeyPath = "C:/Users/Patricio/Google Drive/proyectoYangee/codigoGo/src/github.com/pgmonzon/ServiciosYng/config/keys/app.rsa.pub"
    ExpiraToken   = 10 // en minutos
)

var (
	verifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Inicializar() {
  signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}
