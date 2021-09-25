package main

import (
	"flag"
	"log"
	"web-server/config"
	"web-server/logger"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var command string
	var version int
	flag.StringVar(&command, "command", "up", "Steps to m")
	flag.IntVar(&version, "version", 0, "version to force to")
	flag.Parse()
	log.Println("command is ", command)
	log.Println("version is ", version)

	conf := config.Get().DatabaseConfig
	logger.Infof("migration database config is %v", conf)

	m, err := migrate.New(
		"file://db/migrations",
		conf.DatabaseUrl)
	if err != nil {
		log.Panicf("Failed to connect to database %v", err)
	} else {
		log.Println("Connected to database..")
	}

	switch command {
	case "up":
		log.Println("migrate up")
		err = m.Up()
		if err != nil {
			log.Println("err ", err)
		}
		break
	case "down":
		log.Println("migrate down")
		err = m.Down()
		if err != nil {
			log.Println("err ", err)
		}
		break
	case "force":
		log.Println("migrate force")
		err = m.Force(version)
		if err != nil {
			log.Println("err ", err)
		}
		break
	}
}
