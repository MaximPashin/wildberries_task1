package main

import (
	"fmt"
	"os"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("nats-streaming", "my_publisher")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data, err := os.ReadFile("./message")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Sending")
	sc.Publish("foo", data)
	fmt.Printf("Sended:\n%s\n", string(data))
}
