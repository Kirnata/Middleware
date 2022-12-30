package main

import (
	"github.com/Kirnata/middleware/internal/pkg/app"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
