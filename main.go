package main

import (
	"fmt"
	"github.com/antonioalfa22/go-utils/net"
)

func main() {
	token, code := net.Get("http://192.168.0.15:3000/api/users", nil, true)
	fmt.Println(code)
	fmt.Println(token)
}