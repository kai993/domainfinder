package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		log.Printf("net.Dial: %s\n", err)
		return false, err
	}

	defer conn.Close()

	// 復帰と改行を表す
	conn.Write([]byte(domain + "\r\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}

	return true, nil
}

var marks = map[bool]string{true: "OK", false: "NG"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		exist, err := exists(domain)
		if err != nil {
			fmt.Printf("domain: %s\n", domain)
			log.Fatalln(err)
		}

		if exist {
			// green
			fmt.Printf("\x1b[32m%-30s: %s\x1b[0m\n", domain, marks[exist])
		} else {
			// red
			fmt.Printf("\x1b[31m%-30s: %s\x1b[0m\n", domain, marks[exist])
		}

		time.Sleep(1 + time.Second)
	}
}
