package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		temp := strings.ToLower(strings.Replace(strings.TrimSpace(string(netData)), "get ", "", 1))
		fmt.Printf("GOT: %s\n", temp)
		returns := ""
		valueA := strings.Split(temp, ",")
		rp := regexp.MustCompile("^smtp:")
		for _, value := range valueA {
			fmt.Printf("FOR: %s\n", value)
			if len(rp.FindString(strings.ToLower(value))) > 0 {
				cleanup := strings.Replace(strings.ToLower(value), "smtp:", "", -1)
				returns += "," + cleanup
			}
		}
		fmt.Printf("OBTAINED: %s\n", returns)
		var printme string
		if len(returns) > 0 {
			printme = "200 " + strings.Replace(strings.ToLower(returns), ",", "", 1) + "\n"
		} else {
			printme = "500 no alias found" + "\n"
		}
		fmt.Printf("REPLY: %s\n", printme)
		c.Write([]byte(string(printme)))
		break
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := "127.0.0.1:" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
