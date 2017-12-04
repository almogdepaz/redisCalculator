package main

import (
	"os"
	"fmt"
)

func main(){
	if len(os.Args) != 4 {
		fmt.Println("Usage: redisUrl,inputChannel, outputChannel")
		os.Exit(1)
	}

	redisAdrr := os.Args[1]
	inputChannel := os.Args[2]
	outputChannel := os.Args[3]
	fmt.Println("redis url: " + redisAdrr + ", input channel: " + inputChannel + ", output channel: "+outputChannel)
	handlers := []Handler{Add{},subtract{},multiply{},devide{}}
	run(redisAdrr,inputChannel,outputChannel,10,handlers)
}
