package main

import "log"

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
