package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Addr    string `json:"addr"`
	BaseURL string `json:"base_url"`
	Storage struct {
		Dsn string `json:"dsn"`
	} `json:"storage"`
	Jwt struct {
		SecretKey     string        `json:"secret_key"`
		PeriodInHours time.Duration `json:"period_in_hours"`
	} `json:"jwt"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"redis"`
	Smtp struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		From     string `json:"from"`
	} `json:"smtp"`
}

func NewConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", filepath, err)
	}

	var c *Config

	if err = json.NewDecoder(file).Decode(&c); err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}

	return c, nil
}

func (c *Config) DefaultIfNotSet() {
	c.Addr = set(c.Addr, "localhost:3939")
	c.BaseURL = set(c.BaseURL, "https://localhost:3939")
	c.Storage.Dsn = set(c.Storage.Dsn, "user:pass@localhost:5432/db")
	c.Jwt.SecretKey = set(c.Jwt.SecretKey, "secret")
	c.Redis.Addr = set(c.Redis.Addr, "localhost:6379")
	c.Redis.Password = set(c.Redis.Password, "")
	c.Redis.DB = setInt(c.Redis.DB, 1)
	c.Smtp.Host = set(c.Smtp.Host, "smtp.gmail.com")
	c.Smtp.Port = setInt(c.Smtp.Port, 587)
	c.Smtp.Username = set(c.Smtp.Username, "username")
	c.Smtp.Password = set(c.Smtp.Password, "password")
	c.Smtp.From = set(c.Smtp.From, "email@gmail.com")
}

func set(s, v string) string {
	if s == "" {
		return v
	}
	return s
}

func setInt(s, v int) int {
	if s == 0 {
		return v
	}
	return s
}
