package internal

import (
	"context"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/rashidbazarbaev/urlshortener/base62"
	"github.com/rashidbazarbaev/urlshortener/database"
)

var greeting = "Бот для укорачивания ссылок. Отправьте ему длинную ссылку."
var counters = make(map[int64]int64)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGBOTAPI_KEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, greeting)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			continue
		}

		if !IsReachable(update.Message.Text) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ссылка недействительна")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			continue
		} else {
			url := update.Message.Text
			chatID := update.Message.Chat.ID
			counters[chatID]++
			currentCount := counters[chatID]

			uniqueNumber := chatID*10000 + currentCount
			code := base62.EncodeBase62(uniqueNumber)
			database.DB.Exec(context.Background(),
				"INSERT INTO short_urls (code, original_url) VALUES ($1, $2)",
				code, url,
			)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Ваша ссылка: https://fiercest-derivatively-coreen.ngrok-free.dev/"+code)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}
}
