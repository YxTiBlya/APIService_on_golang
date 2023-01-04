package mailer

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/yxtiblya/internal/cfg"
	"github.com/yxtiblya/internal/models"
	"github.com/yxtiblya/internal/rabbitmq"
	"github.com/yxtiblya/internal/store"
)

// inf handler of mailings
func MailingHandler(q *amqp091.Queue, ch *amqp091.Channel) {
	var mailings []models.Mailing

	for {
		// get all mailings records
		result := store.DB.Find(&mailings)
		if result.Error != nil {
			log.Println(&mailings, result.Error)
			continue
		}

		if len(mailings) != 0 {
			curr_t := time.Now()

			for _, mailing := range mailings {

				// skip mailing if curr time > mailing.end_time or curr time < mailing.star_time
				if mailing.Start_time.After(curr_t) && curr_t.After(mailing.End_time) {
					continue
				}

				json, err := json.Marshal(mailing)
				if err != nil {
					panic(err)
				}

				// send msg to rabbitmq with json of the mailing
				if err := rabbitmq.SendMsg(json, q, ch); err != nil {
					panic(err)
				}

				log.Println("Mailing ID", mailing.ID, "sended to rabbit queue")
			}

		} else {
			log.Println("No mailings")
		}

		// wait some time
		time.Sleep(cfg.GetConfig().MailerExp * time.Second)
	}
}

// inf handler of messages
func MessageHandler(q *amqp091.Queue, ch *amqp091.Channel) {
	// create consume
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println("Failed to register a consumer")
		return
	}

	var mailing models.Mailing
	var contacts []models.Contact
	var messages []models.Message

	// iteration of msgs from rabbit
	for d := range msgs {
		err := json.Unmarshal(d.Body, &mailing)
		if err != nil {
			log.Println(err)
			continue
		}

		// get all records of contacts where tag = ?
		result := store.DB.Where("tag IN ?", strings.Split(mailing.Filters, ",")).Find(&contacts)
		if result.Error != nil {
			panic(result.Error)
		}

		for _, contact := range contacts {

			// checking a record for existence
			result := store.DB.Where("contact_id", contact.ID).Where("mailing_id", mailing.ID).Order("id").Find(&messages)
			if result.Error != nil {
				panic(err)
			}

			// create a new record if not exist
			if len(messages) == 0 {
				curr_t := time.Now()

				message := &models.Message{
					Datetime:   curr_t,
					Status:     "Отправлено",
					Mailing_id: mailing.ID,
					Contact_id: contact.ID,
				}

				// conditional sending of a successful request and data recording
				if result := store.DB.Create(message); result.Error != nil {
					panic(result.Error)
				}

				log.Println("Message Id", message.ID, "created")
			}
		}
	}
}
