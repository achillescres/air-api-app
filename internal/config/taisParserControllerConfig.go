package config

import (
	"path"
	"sync"
	"time"
)

const taisParserControllerConfigFilename = "taisParserController.config.yaml"

type TaisParserControllerConfig struct {
	FTPCheckTimeoutSeconds int    `yaml:"FTPCheckTimeoutSeconds"`
	BucketName             string `yaml:"bucketName"`
	TaisDirPath            string `yaml:"taisDirPath"`
	FileDownloadTLSeconds  int    `yaml:"fileDownloadTLSeconds"`
	FileUploadTLSeconds    int    `yaml:"fileUploadTLSeconds"`

	TaisDirAbsPath  string
	FTPCheckTimeout time.Duration
	FileDownloadTL  time.Duration
	FileUploadTL    time.Duration
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
		taisParserControllerCfgInst.FileDownloadTL = time.Second * time.Duration(taisParserControllerCfgInst.FileDownloadTLSeconds)
		taisParserControllerCfgInst.FileUploadTL = time.Second * time.Duration(taisParserControllerCfgInst.FileUploadTLSeconds)
	})

	return *taisParserControllerCfgInst
}
