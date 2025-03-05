package main

import (
	"demo-go/config"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	config := config.InitSysConfig()
	fmt.Printf("config: %v\n", config)
}
