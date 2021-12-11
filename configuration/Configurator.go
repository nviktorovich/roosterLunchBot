package configuration

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

const ConfigPath = "./configuration/Config.yaml"

type Config struct {
	BotConfig struct {
		API        string `yaml:"api_key"`
		Link       string `yaml:"link"`
		TimeToPoll string `yaml:"time_to_poll"`
		Interval   int    `yaml:"interval"`
	} `yaml:"bot_config"`

	PollConfig struct {
		ChatID              int64    `yaml:"chat_id"`
		Question            string   `yaml:"question"`
		EndPollText         string   `yaml:"end_poll_text"`
		VoteType            string   `yaml:"type"`
		VoteOptions         []string `yaml:"options"`
		IsAnonymous         bool     `yaml:"is_anonymous"`
		CorrectOptionID     int64    `yaml:"сorrect_option_id"`
		AllowMultipleAnswer bool     `yaml:"allows_multiple_answers"`
		OpenPeriod          int      `yaml:"open_period"`
	} `yaml:"poll_config"`
}

func GetConfig() (*Config, error) {
	c := new(Config)
	data, err := os.ReadFile(ConfigPath)

	if err != nil {
		err = errors.New(fmt.Sprintf("error read configuration file, %v", ConfigPath))
		return nil, err
	}

	if err = yaml.Unmarshal(data, &c); err != nil {
		err = errors.New(fmt.Sprintf("error parse configuration file, %v", ConfigPath))
		return nil, err
	}

	return c, nil
}
