package config

import (
	"path"
	"sync"
	"time"
)

const taisParserControllerConfigFilename = "taisParserController.config.yaml"

type TaisParserControllerConfig struct {
	FTPCheckTimeoutSeconds       int    `yaml:"FTPCheckTimeoutSeconds" env-required:"true"`
	BucketName                   string `yaml:"bucketName" env-required:"true"`
	TaisDirPath                  string `yaml:"taisDirPath" env-required:"true"`
	FileDownloadTimeLimitSeconds int    `yaml:"fileDownloadTimeLimit" env-required:"true"`

	TaisDirAbsPath        string
	FTPCheckTimeout       time.Duration
	FileDownloadTimeLimit time.Duration
}

var (
	taisParserControllerCfgInst     = &TaisParserControllerConfig{}
	loadTaisParserControllerCfgOnce = &sync.Once{}
)

func TaisParserController() TaisParserControllerConfig {
	loadTaisParserControllerCfgOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, taisParserControllerConfigFilename), taisParserControllerCfgInst)
		taisParserControllerCfgInst.FTPCheckTimeout = time.Second * time.Duration(taisParserControllerCfgInst.FTPCheckTimeoutSeconds)
		taisParserControllerCfgInst.TaisDirAbsPath = path.Join(Env().ProjectAbsPath, taisParserControllerCfgInst.TaisDirPath)
		taisParserControllerCfgInst.FileDownloadTimeLimit = time.Second * time.Duration(taisParserControllerCfgInst.FileDownloadTimeLimitSeconds)
	})

	return *taisParserControllerCfgInst
}
