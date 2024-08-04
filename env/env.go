package env

import (
	"os"
	"strconv"
)

var (
	BotToken          string
	ChatId            string
	EnableSheduleTest bool
)

func EnvInit() {
	BotToken = os.Getenv("BOT_TOKEN")
	ChatId = os.Getenv("CHAT_ID")
	EnableSheduleTest, _ = strconv.ParseBool(os.Getenv("ENABLE_SHEDULE_TEST"))
}
