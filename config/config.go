package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Specification struct {
	LOG_LEVEL log.Level `default:"info" log:"log.SetLevel()"`
}

func GetConfig() (Specification, error) {
	var s Specification
	err := envconfig.Process("ravencli", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s, nil
}
