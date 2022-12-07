package config

import (
	"path"
	"sync"
)

const middlewareConfigFilename = "middleware.config.yaml"

type MiddlewareConfig struct {
	UserIdCtxKey        string `yaml:"userIdCtxKey" env-default:"userId"`
	AuthorizationHeader string `yaml:"authorizationHeader" en-default:"Authorization"`
}

var (
	middlewareCfgInst     = &MiddlewareConfig{}
	loadMiddlewareCfgOnce = sync.Once{}
)

func Middleware() MiddlewareConfig {
	loadMiddlewareCfgOnce.Do(func() {
		readConfig(path.Join(Env().ConfigAbsPath, middlewareConfigFilename), middlewareCfgInst)
	})

	return *middlewareCfgInst
}
