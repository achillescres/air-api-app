package main

import (
	"api-app/internal/application"
	"api-app/pkg/mylogging"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	mylogging.ConfigureLogrus()
}

func main() {
	ctx := context.Background()
	log.Infoln("Let's go!")
	app, err := application.NewApp(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("fatal creating new app: %s\n", err.Error()))
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatalf("fatal on or while running app: %s\n", err.Error())
	}
}
