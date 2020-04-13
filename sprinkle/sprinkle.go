package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
}

func main() {
	var s *bufio.Scanner
	rand.Seed(time.Now().UTC().UnixNano())

	// ファイル
	if len(os.Args) > 1 {
		filename := os.Args[1]
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		s = bufio.NewScanner(f)
	} else {
		// 標準入力
		s = bufio.NewScanner(os.Stdin)
	}

	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
