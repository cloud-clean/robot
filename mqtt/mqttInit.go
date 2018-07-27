package mqtt

import (
	"sync"
	"fmt"
)

var mc *MqttClient
var once sync.Once

func MqttInit(){
	once.Do(func() {
		mc = NewMqttClient("lot1","cloudhai")
	})
}

func SendString(topic,msg string){
	mc.Client.Publish(topic,1,false,msg)
}

func Send(topic string,msg []byte)error{
	token := mc.Client.Publish(topic,1,false,msg)
	if token.Error() != nil{
		return token.Error()
	}else{
		return nil
	}
}

func Status(){
	if mc.Client.IsConnected(){
		fmt.Println("连接")
	}else{
		fmt.Println("断开连接")
	}
}
