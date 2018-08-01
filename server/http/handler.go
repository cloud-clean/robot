package http

import (
	"robot/entity"
	"robot/mqtt"
	"encoding/json"
	"fmt"
	"robot/cron"
	"time"
	"strconv"
	"robot/store/boltDb"
)

func PinHandler(params *Params)*WebResult{

	pos := params.Get("pos")
	status := params.Get("status")
	msg := entity.SwitchEntity{Position:pos,Status:status}
	data,_ := json.Marshal(msg)
	start:=time.Now()
	err := mqtt.Send("lot",data)
	dur := time.Now().Sub(start)
	log.Infof("after send time:%d",dur.Nanoseconds())
	if err != nil{
		fmt.Println(err.Error())
		return NewResult(FAIL,"error")
	}else {

		return NewResult(SUCCESS,"success")
	}
}

func CronJob(params *Params) *WebResult{
	timestr := params.Get("time")
	timeInt,err := strconv.ParseInt(timestr,10,64)
	if err != nil{
		return NewResult(FAIL,err.Error())
	}
	spec := time.Unix(timeInt,0)
	pos := params.Get("pos")
	status := params.Get("status")
	key := params.Get("name")
	err = cron.AddJob(key,spec,func(){
		mqtt.SendPos(pos,status)
	})
	if err != nil{
		return NewResult(FAIL,err.Error())
	}
	return NewResult(SUCCESS,"success")
}

func GetFromDb(params *Params) *WebResult{
	key := params.Get("key")
	if key == ""{
		return NewResult(FAIL,"ke is nil")
	}
	value := boltDb.GetString(key)
	return NewResult(SUCCESS,entity.SwitchEntity{Position:key,Status:value})
}

func SetToDb(params *Params) *WebResult{
	key := params.Get("key")
	value := params.Get("value")
	if key == "" || value == ""{
		return NewResult(FAIL,"key or value is nil")
	}
	if boltDb.PutString(key,value){
		return NewResult(SUCCESS,"success")
	}
	return NewResult(FAIL,"store to db fail")
}
