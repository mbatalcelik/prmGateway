package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type (
	// Config provides the system configuration.
	Config struct {
		MyApp       MyApp
		Logging     Logging
		BusinessApp BusinessApp
	}
	// Logging provides the logging configuration.
	Logging struct {
		Debug  bool `envconfig:"LOGS_DEBUG"`
		Trace  bool `envconfig:"LOGS_TRACE"`
		Color  bool `envconfig:"LOGS_COLOR"`
		Pretty bool `envconfig:"LOGS_PRETTY"`
		Text   bool `envconfig:"LOGS_TEXT"`
	}
	// Server provides the server configuration.
	MyApp struct {
		Port string `envconfig:"MYAPP_PORT" required:"true"`
	}
	BusinessApp struct {
		Url      string `envconfig:"BUSINESSAPP_URL" default:"https://jsonplaceholder.typicode.com/"`
		TestPath string `envconfig:"BUSINESSAPP_TESTPATH" default:"posts"`
	}
)

func GetFromEnv() (Config, error) {
	var c Config
	err := envconfig.Process("PRMGATECONF", &c)
	if err != nil {
		log.Fatal("conf hata" + err.Error())
	}
	return c, err
}
