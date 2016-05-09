package srygoclient

import (
	"fmt"
)

type SryGoClient struct {
	Username string
	Password string
	Host     string
	Token    string
}

func NewSryGoClient(username, password) *SryGoClient {
	return nil
}

func Ping(client *SryGoClient) bool {
	return false
}

func Auth(client *SryGoClient) bool {
	return true
}

func Login(client *SryGoClient) bool {
	return Auth(client)
}
