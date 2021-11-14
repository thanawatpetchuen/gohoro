package main

import (
	"log"
	"os"
	"strings"
)

func Write(fileName string, data string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		pathBoom := strings.Split(fileName, "/")
		path := strings.Join(pathBoom[:len(pathBoom)-1], "")
		os.MkdirAll(path, 0700) // Create your file
	}
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

	}()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(data + "\n")); err != nil {
		log.Fatal(err)
	}
}
