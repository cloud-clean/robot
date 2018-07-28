package main

import (
	"robot/mqtt"
	"robot/server/http"
	"robot/cron"
)

func main(){
	go mqtt.MqttInit()
	go http.Start(8052)
	go cron.CronInit()
	select {

	}
}
