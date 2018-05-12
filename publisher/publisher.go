package main

import (
	"os"
	"fmt"
	"encoding/json"
	"redis_calculator/commons"
)



func main() {
	 //The user name is the only argument, so we'll grab that here
	if len(os.Args) != 4 {
		fmt.Println("Usage: goredchat redisUrl")
		os.Exit(1)
	}

	redisAdrr := os.Args[1]
	inputChannel := os.Args[2]
	outputChannel := os.Args[3]
	fmt.Println("redis url: " + redisAdrr + "input channel: " + inputChannel + "output channel: " + outputChannel)
	conn := commons.GetConnection(redisAdrr)

	defer conn.Close()

	fmt.Println("connected to redis server")

	calc := commons.Input{
	}
	calc.X = "3"
	calc.Y = "8"
	calc.Action = "+"
	res, _ := json.Marshal(calc)
	conn.Do("PUBLISH", inputChannel, string(res))

	calc.X = "3"
	calc.Y = "8"
	calc.Action = "-"
	res, _ = json.Marshal(calc)

	conn.Do("PUBLISH", inputChannel, string(res))

	calc.X = "3"
	calc.Y = "8"
	calc.Action = "*"
	res, _ = json.Marshal(calc)
	conn.Do("PUBLISH", inputChannel, string(res))

	calc.X = "3"
	calc.Y = "8"
	calc.Action = "/"
	res, _ = json.Marshal(calc)
	conn.Do("PUBLISH", inputChannel, string(res))

	calc.X = "gdadf"
	calc.Y = "8"
	calc.Action = "+"
	res, _ = json.Marshal(calc)
	conn.Do("PUBLISH", inputChannel, string(res))

	calc.X = "3"
	calc.Y = "0"
	calc.Action = "/"
	res, _ = json.Marshal(calc)

	conn.Do("PUBLISH", inputChannel, string(res))
}






