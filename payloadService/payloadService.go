package payloadservice

import (
	"fmt"
	nats "practice/nats"
	"regexp"
	"strconv"
	"strings"
	"time"

	natsOuter "github.com/nats-io/nats.go"
)

//Слушаем канал натс и добавляем/удаляем горутины. так же генерить id. считать кол-во запущенных сервисов и аптайм. данные эти отправлять по натсу контроллеру.

var instanceCounter int
var InstanceSignal chan int

func InstanceHandler() {
	InstanceSignal = make(chan int)

	nats.Nc.Subscribe("instance", func(m *natsOuter.Msg) {
		msg := string(m.Data)
		match, _ := regexp.MatchString(`^del\s\d+`, msg)
		if msg == "add" {
			instanceCounter++
			id := instanceCounter
			go instanceCreate(id)
			fmt.Println("succesfully started instance")
		} else if match {
			msgSplit := strings.Split(msg, " ")
			deletedId, _ := strconv.Atoi(msgSplit[len(msgSplit)-1])
			// ???
			fmt.Println(deletedId, " deletedId")
			InstanceSignal <- deletedId
		} else {
			nats.Nc.Publish("uptime", []byte("unknown command has been sent to payloadService"))
		}
	})
}

func instanceCreate(id int) {
	uptime := time.Now().String()
	instance := Instance{id, uptime}

	nats.Nc.Publish("uptime", []byte(fmt.Sprintf("%+v. active instances - %d", instance, instanceCounter)))

	//Тут функция должна делать полезную работу

	//???
	for {
		select {
		case recievedId := <-InstanceSignal:
			fmt.Println(recievedId, " recievedId")
			fmt.Println(id, " id")
			if recievedId == id {
				fmt.Println("succesfully deleted instance")
				instanceCounter--
				return
			}
		default:
			continue
		}
	}
}
