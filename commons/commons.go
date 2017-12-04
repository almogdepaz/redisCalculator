package commons

import (
	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
	"fmt"
	"os"
)

type Input struct {
	X  		   string
	Y    	   string
	Action     string
}


func GetConnection(redisAdrr string) redis.Conn {
	conn, err := redisurl.ConnectToURL(redisAdrr)
	CheckError(err)
	return conn
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
