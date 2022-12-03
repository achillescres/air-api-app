package main

import (
	"context"
	"fmt"
	"github.com/achillescres/saina-api/internal/application"
	"github.com/achillescres/saina-api/pkg/mylogging"
	log "github.com/sirupsen/logrus"
)

func init() {
	mylogging.ConfigureLogrus()
}

func main() {
	log.Infoln("Let's go!")
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
