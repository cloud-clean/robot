package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"robot/cron"
)

func main(){
	cron.CronInit()
	cron.AddJob("test",time.Now().Add(time.Second*10),func(){
		fmt.Println("job is running....")
	})
	select{

	}


}

func server(){
	svr,err := net.ResolveTCPAddr("tcp","127.0.0.1:8322")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	listen,err := net.ListenTCP("tcp",svr)
	if err != nil{
		fmt.Println(err)
	}
	defer listen.Close()
	conn,err := listen.Accept()
	defer conn.Close()
	var buffer [2048]byte
	conn.Read(buffer[:])
	request := string(buffer[:])
	fmt.Println(strings.Split(request,"\r\n")[0])
	fmt.Println(string(buffer[:]))
}
