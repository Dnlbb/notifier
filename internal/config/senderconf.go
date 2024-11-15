package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	token = "BOT_TOKEN"
	id    = "ID"
)

type SenderConf struct {
	token string
	id    string
}

type Sender interface {
	Token() string
	ID() int64
}

func NewSenderConf() (Sender, error) {
	token := os.Getenv(token)
	if len(token) == 0 {
		return nil, fmt.Errorf("token not found in environment variable %s", token)
	}

	ID := os.Getenv(id)
	if len(ID) == 0 {
		return nil, fmt.Errorf("id not found in environment variable %s", id)
	}

	return SenderConf{token, ID}, nil
}

func (s SenderConf) Token() string {
	return s.token
}

func (s SenderConf) ID() int64 {
	intID, err := strconv.Atoi(s.id)
	if err != nil {
		log.Fatal("error converting id to int")
	}
	return int64(intID)
}
