package main

import (
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/yxtiblya/internal/cfg"
	"github.com/yxtiblya/internal/mailer"
	"github.com/yxtiblya/internal/rabbitmq"
	"github.com/yxtiblya/internal/store"
)

func main() {
	// parsing toml file to config
	config := cfg.NewConfig()
	if _, err := toml.DecodeFile("configs/config.toml", config); err != nil {
		log.Println("failed to decode toml file")
		return
	}
	cfg.ChangeConfig(config)

	// logger writer in file and console
	f, _ := os.Create("logger/mailer.log")
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	log.Println("Start mailer service")

	// connection to db
	_, err := store.NewDB(config.DatabaseURL)
	if err != nil {
		log.Println("failed to connect database")
		return
	}

	// connection to rabbitmq
	ch, err := rabbitmq.NewChannel()
	if err != nil {
		log.Println(err)
		return
	}
	defer ch.Close()

	// create the queue
	q, err := ch.QueueDeclare("mailings", true, false, false, false, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var forever chan struct{}

	go mailer.MailingHandler(&q, ch)
	go mailer.MessageHandler(&q, ch)

	<-forever
}
