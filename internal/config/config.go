package config

import (
	"api-app/pkg/gconfig"
	log "github.com/sirupsen/logrus"
)

type Config interface{}

func readConfig(cfgPath string, cfgInst Config) {
	err := gconfig.ReadConfig(cfgPath, cfgInst)
	if err != nil {
		log.Fatalf("fatal reading %s: %s", cfgPath, err.Error())
	}
}
