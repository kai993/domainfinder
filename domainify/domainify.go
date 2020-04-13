package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var inputTld = flag.String("tld", "", "input top level domain(.com, .net...)")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())

		if len(text) == 0 {
			continue
		}

		var newText []rune
		for _, r := range text {
			// スペースは-に置換
			if unicode.IsSpace(r) {
				r = '-'
			}

			// 許可されていない文字は削除
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}

			newText = append(newText, r)
		}

		// select top level domain
		var tld string
		if len(*inputTld) == 0 {
			tld = tlds[rand.Intn(len(tlds))]
		} else {
			tld = *inputTld
		}

		fmt.Println(string(newText) + "." + tld)
	}
}
