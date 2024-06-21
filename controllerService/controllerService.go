package controllerservice

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"practice/nats"
	"regexp"
	"strconv"
	"strings"

	natsOuter "github.com/nats-io/nats.go"
)

func OperateInstances() {
	var instancesHistory []string
	nats.Nc.Subscribe("uptime", func(m *natsOuter.Msg) {
		instancesHistory = append(instancesHistory, string(m.Data))
		fmt.Println(string(m.Data))
	})

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		match, _ := regexp.MatchString(`^del\s\d+`, input)
		switch {
		case input == "add":
			resp, _ := http.Get("http://localhost:8080/addInstance")
			defer resp.Body.Close()
		case match:
			msgSplit := strings.Split(input, " ")
			deletedId, _ := strconv.Atoi(msgSplit[len(msgSplit)-1])
			id := Id{deletedId}
			json, _ := json.Marshal(id)
			resp, _ := http.Post("http://localhost:8080/deleteInstance", "application/json", bytes.NewReader(json))
			defer resp.Body.Close()
		case input == "status":
			fmt.Println(instancesHistory)
		default:
			fmt.Println("unknown command. Use add or del INSTANCE_ID")
		}
	}
}
