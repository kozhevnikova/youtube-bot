package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var telegramToken string

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR >", err)
	}
	fmt.Fprintln(os.Stdout, "Authorized on account", bot.Self.UserName)

	conf := tgbotapi.NewUpdate(0)
	conf.Timeout = 60

	updates, err := bot.GetUpdatesChan(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR >", err)
	}

	for update := range updates {
		chat := update.Message.Chat.ID
		command := update.Message.Command()

		switch command {
		case "start":
			_, err := bot.Send(
				tgbotapi.NewMessage(chat,
					"Welcome to SaveAudioBot. I can convert video to audio (mp3). Just give me a link from youtube."))
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR >", err)
			}

		case "":
			id := strings.TrimPrefix(
				update.Message.Text, "https://www.youtube.com/watch?v=")
			fmt.Println(id)

			_, err := bot.Send(
				tgbotapi.NewMessage(chat,
					"Your request has been received. It can take a minute to process."))
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR >", err)
			}

			video, err := getVideoStream(id)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR get video >", err)
				_, err := bot.Send(
					tgbotapi.NewMessage(chat,
						"Video is unavailable. Check the link."))
				if err != nil {
					fmt.Fprintln(os.Stderr, "ERROR >", err)
				}
			} else {
				_, err = bot.Send(
					tgbotapi.NewMessage(chat,
						"Downloading. Please, wait."))
				if err != nil {
					fmt.Fprintln(os.Stderr, "ERROR >", err)
				}

				err = downloadVideo(video, id)
				if err != nil {
					fmt.Fprintln(os.Stderr, "ERROR download video >", err)
					_, err = bot.Send(
						tgbotapi.NewMessage(chat,
							"This video is unavailable. Check the link"))
					if err != nil {
						fmt.Fprintln(os.Stderr, "ERROR >", err)
					}
				} else {
					reply, err := getAudio(id)
					if err != nil {
						fmt.Fprintln(os.Stderr, "ERROR get audio file >", err)
						_, err = bot.Send(
							tgbotapi.NewMessage(chat,
								"Can't get audio. Please, send me link again"))
						if err != nil {
							fmt.Fprintln(os.Stderr, "ERROR >", err)
						}
					} else {
						_, err := bot.Send(
							tgbotapi.NewMessage(chat,
								"Sending audio"))
						if err != nil {
							fmt.Fprintln(os.Stderr, "ERROR >", err)
						}

						_, err = bot.Send(tgbotapi.NewAudioUpload(chat, reply))
						if err != nil {
							fmt.Fprintln(os.Stderr, "ERROR audio sending >", err)
						}
					}
				}
			}
		default:
			_, err = bot.Send(
				tgbotapi.NewMessage(chat,
					"Can't understand your instruction. Use /start for information"))
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR >", err)
			}
		}
	}
}
