package main

import (
	"fmt"
	"visa"
)

func main() {
	conn, err:= visa.Open("tcpip0::t6290e-c00119.local::hislip1", 1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	resp := conn.Query("*IDN?")
	fmt.Println("idn", resp)
}
