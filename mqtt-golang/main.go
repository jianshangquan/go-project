package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	options := mqtt.NewClientOptions().AddBroker("tls://133866ccfd484dfabe8e1799e4c34719.s1.eu.hivemq.cloud:8883")

	options.SetClientID("go_mqtt_client")
	options.SetUsername("jianshangquan")
	options.SetPassword("Jsq123123")
	client := mqtt.NewClient(options)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	text := "Hello from Golang MQTT!"
	token := client.Publish("test/topic", 0, false, text)
	token.Wait()

	fmt.Println("Published message:", text)
	// client.Disconnect(250)
}
