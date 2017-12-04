package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
	"strings"
	"encoding/json"
	"redisCalculator/commons"
)


func run(redisAdrr string, inputChannel string, outputChannel string,bufferSize int,handlers []Handler) {

	//map action to handler
	handlerMap := make(map[string]Handler)
	var keys []string
	fmt.Print(strings.Join(keys, ","))
	for i := 0; i < len(handlers); i += 1 {
		handlerMap[handlers[i].applysTo()] = handlers[i]
		keys = append(keys, handlers[i].applysTo())
	}
	fmt.Println("handles: " + strings.Join(keys, ","))

	conn := commons.GetConnection(redisAdrr)
	defer conn.Close()

	fmt.Println("connected to redis server")

	//goRouting that listens to redis channel
	subChan := make(chan string, bufferSize)
	go func() {
		subconn, err := redisurl.ConnectToURL(redisAdrr)
		commons.CheckError(err)
		defer subconn.Close()
		psc := redis.PubSubConn{Conn: subconn}
		psc.Subscribe(inputChannel)
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				subChan <- string(v.Data)
			case error:
				return
			}
		}
	}()

	for {
		select {
		case msg := <-subChan:
			calc := commons.Input{}
			json.Unmarshal([]byte(msg), &calc)
			fmt.Println("input is: " + msg)
			res := handlerMap[calc.Action].apply(calc)
			fmt.Println("calc result was: " + res)
			conn.Do("PUBLISH", outputChannel, res)
		}
	}

}


