package config

import (
  "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var environments = map[string]string{
  "produccion": "C:/Users/Patricio/Google Drive/proyectoYangee/codigoGo/src/github.com/pgmonzon/ServiciosYng/config/prod.json",
  "desarrollo": "C:/Users/Patricio/Google Drive/proyectoYangee/codigoGo/src/github.com/pgmonzon/ServiciosYng/config/desa.json",
}

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings Settings = Settings{}
var env = "desarrollo"

func Inicializar() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: No se encontró entorno, se seteó desarrollo")
		env = "desarrollo"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error: No se pudo leer el config", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error: No se pudo parsear el config", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Inicializar()
	}
	return settings
}
