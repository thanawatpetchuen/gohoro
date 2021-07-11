package main

import (
	"flag"
	"log"
	"sync"
)

func main() {
	iter := flag.Int("i", 0, "Number of iteration")
	flag.Parse()

	app := NewApp()

	// exitCh := make(chan os.Signal, 1)
	// signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)

	if *iter > 0 {
		log.Println("[Iteration mode]")
		for i := 1; i <= *iter; i++ {
			app.Start(i, nil)
		}
	} else {
		log.Println("[Multithread mode]")
		workers := app.GetWorkers()
		wg := sync.WaitGroup{}

		for i := 1; i <= workers; i++ {
			wg.Add(1)
			go app.Start(i, &wg)
		}

		wg.Wait()
		log.Println("All workers done.")
	}

	// <-exitCh
	log.Println("Shutting down...")
	app.Stop()
}
