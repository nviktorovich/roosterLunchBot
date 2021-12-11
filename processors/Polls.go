package processors

import (
	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nviktorovich/roosterLunchBot/configuration"
)

func GetPoll(p configuration.Config) *telebot.SendPollConfig {
	poll := telebot.SendPollConfig{
		BaseChat:              telebot.BaseChat{ChatID: p.PollConfig.ChatID},
		Question:              p.PollConfig.Question,
		Options:               p.PollConfig.VoteOptions,
		IsAnonymous:           p.PollConfig.IsAnonymous,
		Type:                  p.PollConfig.VoteType,
		AllowsMultipleAnswers: p.PollConfig.AllowMultipleAnswer,
		CorrectOptionID:       p.PollConfig.CorrectOptionID,
		Explanation:           "",
		ExplanationParseMode:  "",
		ExplanationEntities:   nil,
		OpenPeriod:            p.PollConfig.OpenPeriod,
		CloseDate:             0,
		IsClosed:              false,
	}
	return &poll
}
