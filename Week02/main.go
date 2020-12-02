package main

import (
	"fmt"
	"os"
	"strconv"

	"Week02/service"
)

func main() {
	s := service.AccountService{}
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		input = 0
	}
	us, err := s.GetUser(input)
	if err != nil {
		fmt.Printf("%+v\r\n", err)
	}
	fmt.Println(us)
}
