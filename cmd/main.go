package main

import (
	"api-app/internal/application"
	"api-app/internal/config"
	"api-app/pkg/mylogging"
	"api-app/pkg/security/passlib"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	mylogging.ConfigureLogrus()
}

func main() {
	log.Infoln("Let's go!")

	passHashSalt := config.Env().PasswordHashSalt
	passlib.Init(passHashSalt)

	ctx := context.Background()
	log.Infoln("Creating app...")
	app, err := application.NewApp(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("fatal creating new app: %s\n", err.Error()))
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatalf("fatal on or while running app: %s\n", err.Error())
	}
}
