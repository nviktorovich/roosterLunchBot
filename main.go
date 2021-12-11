package main

import (
	"fmt"
	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nviktorovich/roosterLunchBot/configuration"
	"github.com/nviktorovich/roosterLunchBot/processors"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	cfg, err := configuration.GetConfig()
	e := new(processors.LogRow)

	if err != nil {
		e.Log = err
		e.Renew()
		logrus.Fatalf("%v", e.Show())
	}
	bot, err := telebot.NewBotAPI(cfg.BotConfig.API)

	if err != nil {
		e.Log = err
		e.Renew()
		logrus.Fatalln(e.Show())
	}

	if _, err = bot.Send(processors.GetPoll(*cfg)); err != nil {
		e.Log = err
		e.Renew()
		logrus.Error(e.Show())
	}
	fmt.Println(cfg.BotConfig.Interval)

	c := new(processors.CloseMessage)
	c.Info = cfg.PollConfig.EndPollText

	func() {
		for i := 0; i < cfg.BotConfig.Interval; i++ {
			time.Sleep(time.Minute)
		}
	}()

	c.Renew()
	msg := telebot.NewMessage(cfg.PollConfig.ChatID, c.Show())
	if _, err = bot.Send(msg); err != nil {
		e.Log = err
		e.Renew()
		logrus.Error(e.Show())
	}

}
