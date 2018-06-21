package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoURL = "mongodb://192.168.1.113/data_manage"

func main() {
	session, err := mgo.DialWithTimeout(mongoURL, 10*time.Second)
	if nil != err {
		fmt.Println(err)
	}
	go func() {
		for {
			time.Sleep(10 * time.Second)
			session.Refresh()
		}
	}()
	c := session.DB("").C("person")
	condTmp := bson.M{}
	var result []interface{}

	err = c.Find(condTmp).Select("baseinfo.idcard").Distinct("baseinfo.idcard", &result)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(len(result))
	}

}
