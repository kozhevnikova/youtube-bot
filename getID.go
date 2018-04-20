package main

import (
	"regexp"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func getID(update tgbotapi.Update, url string) (string, error) {
	var id string
	m, err := regexp.MatchString("m.youtube", url)
	if err != nil {
		return "", err
	}
	w, err := regexp.MatchString("www.youtube", url)
	if err != nil {
		return "", err
	}
	if m == true {
		id = strings.TrimPrefix(
			update.Message.Text, "https://m.youtube.com/watch?v=")

	} else if w == true {
		id = strings.TrimPrefix(
			update.Message.Text, "https://www.youtube.com/watch?v=")
	} else {
		id = ""
	}
	return id, nil
}
