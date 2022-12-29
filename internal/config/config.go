package config

import (
	"github.com/achillescres/saina-api/pkg/gconfig"
	log "github.com/sirupsen/logrus"
)

func readConfig(cfgPath string, cfgInst gconfig.Config) {
	log.Infof("reading %s\n", cfgPath)
	err := gconfig.ReadConfig(cfgPath, cfgInst)
	if err != nil {
		//log.Fatalf("fatal reading config: %s\n", err)
	}
	//log.Infof("successfully read %s\n", cfgPath)
}
