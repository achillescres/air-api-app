package gconfig

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config interface{}

func ReadConfig(path string, inst Config) error {
	if err := cleanenv.ReadConfig(path, inst); err != nil {
		return err
	}

	return nil
}

func ReadEnv(filename string, inst Config) error {
	err := godotenv.Load(filename)
	if err != nil {
		return err
	}
	err = cleanenv.ReadEnv(inst)
	if err != nil {
		return err
	}

	return nil
}
