package rfmt

import (
	"strings"
)

type Configuration struct {
	root  string
	api   string
	delim string
}

var conf = &Configuration{
	root:  "/",
	api:   "api/",
	delim: "/",
}

func Api() string {
	return conf.api
}

func Root() string {
	return conf.root
}

func Delim() string {
	return conf.delim
}

func SetRoot(root string) {
	conf.root = root
}

func SetApi(api string) {
	conf.api = api
}

func SetDelim(delim string) {
	conf.delim = delim
}

func Conf() Configuration {
	return *conf
}

func SetConf(newConf *Configuration) {
	conf.api = newConf.api
	conf.root = newConf.root
}

func Join(els ...string) string {
	return strings.Join(els, Delim())
}

func JoinRoot(els ...string) string {
	return Root() + strings.Join(els, conf.delim)
}

func JoinApi(els ...string) string {
	return JoinRoot(Api()) + strings.Join(els, Delim())
}
