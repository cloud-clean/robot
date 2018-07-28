package boltDb

import (
	"github.com/boltdb/bolt"
	"robot/common/logger"
	"os/signal"
	"syscall"
	"os"
	"fmt"
)

var log = logger.NewLog()
var db *bolt.DB = nil
var BUCKET_NAME = []byte("lot")

func init(){
	db,_ = bolt.Open("robot.db",0700,nil)
	//if err != nil{
	//	log.Errorf("init db fail error:%s",err.Error())
	//}
	go func(){
		ss := make(chan os.Signal,1)
		signal.Notify(ss,syscall.SIGEMT,syscall.SIGINT)
		for{
			select {
			case <-ss:
				fmt.Println("program quit")
				db.Close()
				return
			}
		}
	}()
}

func Put(key ,value []byte)bool{
	err := db.Update(func(tx *bolt.Tx) error{
		b,err := tx.CreateBucketIfNotExists(BUCKET_NAME)
		if err != nil{
			return err
		}
		return b.Put(key,value)
	})
	if err != nil{
		return false
	}else{
		return true;
	}
}

func Get(key []byte) []byte{
	var v []byte
	err := db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket(BUCKET_NAME)
		v = b.Get(key)
		return nil
	})
	if err != nil{
		return nil
	}else{
		return v
	}
}

func PutString(key,value string)bool{
	return Put([]byte(key),[]byte(value))
}

func GetString(key string) string{
	value := Get([]byte(key))
	if value == nil{
		return ""
	}
	return string(value)
}


