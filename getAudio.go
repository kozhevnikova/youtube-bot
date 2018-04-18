package main

import (
	"io/ioutil"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func getAudio(id string) (tgbotapi.FileBytes, error) {
	f := tgbotapi.FileBytes{}

	audio, err := ioutil.ReadFile(id + ".mp3")
	if err != nil {
		return f, err
	}

	f = tgbotapi.FileBytes{
		Name:  id + ".mp3",
		Bytes: audio,
	}
	return f, nil
}
