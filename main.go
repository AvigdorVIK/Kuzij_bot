package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5596286474:AAGOvnUuDZS6wmYwCR41tfeU4aOfz_z5mkM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			continue
		}


			command := strings.Split(update.Massage.Text, sep :" ")

			

			switch command [0]{
			case "ADD":
				bot.Send(tgbotapi.NewMessage(update.Message.MigrateToChatID, text:"Valuta dobavlena"))
			case"SUB":
				bot.Send(tgbotapi.NewMessage(update.Message.MigrateToChatID, text:"Valuta sub"))
			case"DEL":
				bot.Send(tgbotapi.NewMessage(update.Message.MigrateToChatID, text:"Valuta del"))
			case"SHOW":
				bot.Send(tgbotapi.NewMessage(update.Message.MigrateToChatID, text:"Valuta show"))
			default:
				bot.Send(tgbotapi.NewMessage(update.Message.MigrateToChatID, text:"Valuta ne naidena"))
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, command[0])

			bot.Send(msg)
		
	}

}
