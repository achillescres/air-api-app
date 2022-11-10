package rfmt

import (
	"path"
)

type Configuration struct {
	root string
	api  string
}

var conf = &Configuration{
	root: "/",
	api:  "/api",
}

func Conf() Configuration {
	return *conf
}

func SetConf(newConf *Configuration) {
	conf.api = newConf.api
	conf.root = newConf.root
}

func JoinRoute(els ...string) string {
	return path.Join(els...)
}

func JoinApiRoute(els ...string) string {
	return path.Join(conf.root, conf.api, path.Join(els...))
}
