package http

import (
	"robot/entity"
	"robot/mqtt"
	"encoding/json"
	"fmt"
)

func PinHandler(params *Params)*WebResult{
	pos := params.Get("pos")
	status := params.Get("status")
	msg := entity.SwitchEntity{Position:pos,Status:status}
	data,_ := json.Marshal(msg)
	err := mqtt.Send("lot",data)
	if err != nil{
		fmt.Println(err.Error())
		return NewResult(0,"error")
	}else {

		return NewResult(1,"success")
	}

}
