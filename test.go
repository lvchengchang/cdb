package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("cdb.dat")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		a := strings.Split(line, "gocdb:")
		for _, v := range a {

		}

		if err != nil || io.EOF == err {
			break
		}
	}
}
