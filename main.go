package main

import (
	"fmt"

	"github.com/chesiver/gomock-playground/clients"
)

func InvokeClient(client clients.Client) {
	client.GetData()
}

func main() {
	fmt.Println("Start!")
}
