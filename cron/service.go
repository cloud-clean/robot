package cron

import (
	"github.com/robfig/cron"
	"time"
	"errors"
	"fmt"
)

var c *cron.Cron
var jobMap = make(map[string] time.Time)
func CronInit(){
	c = cron.New()
	c.Start()
}

func AddJob(key string,spec time.Time,job func())error{
	if _,ok := jobMap[key];ok{
		return errors.New("job is exists")
	}
	jobMap[key] = spec
	c.AddFunc(toSpec(spec),func(){
		fmt.Println("start run the job "+key)
		job()
		delete(jobMap,key)
	})
	return nil
}


func toSpec(time time.Time) string{
	return time.Format("05 04 15 02 01 ?")
}