package gconfig

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config interface{}

func ReadConfig(path string, inst Config) error {
	if err := cleanenv.ReadConfig(path, inst); err != nil {
		return errors.New(err.Error())
	} else {
		return nil
	}
}
