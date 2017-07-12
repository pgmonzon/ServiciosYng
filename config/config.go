package config

import (
  "crypto/rsa"
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
    PrivKeyPath = "C:/Users/Patricio/Google Drive/proyectoYangee/codigoGo/src/github.com/pgmonzon/ServiciosYng/config/keys_desa/private_key"
    PubKeyPath = "C:/Users/Patricio/Google Drive/proyectoYangee/codigoGo/src/github.com/pgmonzon/ServiciosYng/config/keys_desa/public_key.pub"
    ExpiraToken   = 12 // en cantidad de horas
)

var (
  verifyKey *rsa.PublicKey
  signKey   *rsa.PrivateKey
)

func Inicializar() {
/*
  SignBytes, err := ioutil.ReadFile(privKeyPath)
  if err != nil {
    log.Fatal(err)
  }

  SignKey, err := jwt.ParseRSAPublicKeyFromPEM(SignBytes)
  if err != nil {
    log.Fatal(err)
  }

  VerifyBytes, err := ioutil.ReadFile(pubKeyPath)
  if err != nil {
    log.Fatal(err)
  }

  VerifyKey, err := jwt.ParseRSAPublicKeyFromPEM(VerifyBytes)
  if err != nil {
    log.Fatal(err)
  }
*/
}
