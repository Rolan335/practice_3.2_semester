package nats

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

var Nc *nats.Conn
var err error

func Init() {
	Nc, err = nats.Connect(nats.DefaultURL)
	if err != nil{
		fmt.Println(err)
	}
}

func Connection() *nats.Conn{
	return Nc
}