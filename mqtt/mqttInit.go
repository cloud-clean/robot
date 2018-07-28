package mqtt

import (
	"sync"
	"fmt"
	"robot/entity"
	"encoding/json"
	"errors"
)

var mc *MqttClient
var once sync.Once

func MqttInit(){
	once.Do(func() {
		mc = NewMqttClient("lot2","cloudhai")

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

func SendPos(pos,status string) error{
	fmt.Println(pos+"  "+status)
	if status != "0" && status != "1"{
		return errors.New("value is not 0 or 1")
	}
	msg := entity.SwitchEntity{Position:pos,Status:status}
	data,_ := json.Marshal(msg)
	err := Send("lot",data)
	return err
}

func Status(){
	if mc.Client.IsConnected(){
		fmt.Println("连接")
	}else{
		fmt.Println("断开连接")
	}
}

