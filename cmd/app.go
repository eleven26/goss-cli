package cmd

import (
	"log"
	"os"

	"github.com/eleven26/goss/v3"
)

var app *goss.Goss

func Run() {
	var err error
	app, err = goss.New(goss.WithConfig(&goss.Config{
		Endpoint:  os.Getenv("GOSS_ENDPOINT"),
		AccessKey: os.Getenv("GOSS_ACCESS_KEY"),
		SecretKey: os.Getenv("GOSS_SECRET_KEY"),
		Region:    os.Getenv("GOSS_REGION"),
		Bucket:    os.Getenv("GOSS_BUCKET"),
	}))
	if err != nil {
		log.Fatal(err)
	}

	err = Execute()
	if err != nil {
		log.Fatal(err)
	}
}
