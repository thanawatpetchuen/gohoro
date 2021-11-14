package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	ChromeDriver ChromeDriver `mapstructure:"chromeDriver"`
	HoroSite     HoroSite     `mapstructure:"horoSite"`
	Selenium     Selenium     `mapstructure:"selenium"`
	Workers      int
}

type ChromeDriver struct {
	Path      string `mapstructure:"path"`
	Headless  bool   `mapstructure:"headless"`
	MuteAudio bool   `mapstructure:"muteAudio"`
}

type Selenium struct {
	Path  string `mapstructure:"path"`
	Port  int    `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type HoroSite struct {
	URL string `mapstructure:"url"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	vp := viper.New()
	vp.AddConfigPath(".")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := vp.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
