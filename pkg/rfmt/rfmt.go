package rfmt

import (
	"strings"
)

type configuration struct {
	root  string
	api   string
	delim string
}

var conf = &configuration{
	root:  "/",
	api:   "api/",
	delim: "/",
}

func Conf() configuration {
	return *conf
}

func SetConf(newConf *configuration) {
	conf.api = newConf.api
	conf.root = newConf.root
}

func JoinRoute(els ...string) string {
	return conf.delim + strings.Join(els, conf.delim)
}

func JoinApiRoute(els ...string) string {
	return conf.root + conf.api + strings.Join(els, conf.delim)
}
