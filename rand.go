package main

import (
	"crypto/rand"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main(){
	var sb [4]byte
	rand.Read(sb[:])
	fmt.Println(sb)
	fmt.Printf("%x\n",sb[3])
	fmt.Printf("%x\n",sb[2])
	fmt.Printf("%x\n",sb[1])
	fmt.Printf("%x\n",sb[0])
	fmt.Printf("%x\n",int(sb[0])<<8)
	seed := int64(int64(time.Now().Nanosecond()<<32)|int64(sb[0])<<24|int64(sb[1])<<16|int64(sb[2])<<8|int64(sb[3]))
	fmt.Printf("%x\n",seed)
	//server()


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
