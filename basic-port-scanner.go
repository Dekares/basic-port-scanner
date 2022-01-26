package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanPort(hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	host := os.Args[1]
	fmt.Println("Port Scanning")

	for i := 1; i < 65536; i++ {
		open := scanPort(host, i)
		if open {
			fmt.Println("Port Open: ", i)
		} else {
			fmt.Println("Port Close: ", i)
		}

	}

}
