package mqtt

import (
	"sync"
	"fmt"
	"robot/entity"
	"encoding/json"
	"errors"
	"github.com/eclipse/paho.mqtt.golang"
	"robot/store/boltDb"
)

var mc *MqttClient
var once sync.Once

const TOPIC_LOT  = "lot_data"
const mqtt_user = "lot_server"
const mqtt_password = "Kiw28&4292si"

func MqttInit(){
	once.Do(func() {
		mc = NewMqttClient(mqtt_user,mqtt_password)
		mc.Subscribe(TOPIC_LOT, func(client mqtt.Client, message mqtt.Message) {
			log.Infof("get msg from lot_data %s",string(message.Payload()))
			var msg entity.SwitchEntity
			err := json.Unmarshal(message.Payload(),&msg)
			if err != nil{
				log.Errorf("unmarshal fail msg:%s",string(message.Payload()))
			}
			if !boltDb.PutString(msg.Position,msg.Status){
				log.Error("save msg to db fail")
			}
		})

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

