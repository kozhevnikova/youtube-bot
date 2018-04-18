package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.StringVar(&telegramToken, "token", "", "Telegram bot token")
	flag.Parse()

	if telegramToken == "" {
		fmt.Fprintln(os.Stderr, "ERROR > Token not found")
	}
}
