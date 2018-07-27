package main

import (
	"robot/mqtt"
	"robot/server/http"
)

func main(){
	go mqtt.MqttInit()
	go http.Start(8052)
	select {

	}
}
