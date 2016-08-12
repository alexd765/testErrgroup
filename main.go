package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	//	exampleWaitgroup()
	exampleErrgroup()
}

func exampleWaitgroup() {
	var wg sync.WaitGroup

	log.Println("example WaitGroup started")

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
	}

	wg.Wait()
	log.Println("example WaitGroup finished")

}

func exampleErrgroup() {
	var eg errgroup.Group

	log.Println("example ErrGroup started")

	for i := 0; i < 4; i++ {
		eg.Go(func() error {
			time.Sleep(1 * time.Second)
			return errors.New("we got an error")
		})
	}

	err := eg.Wait()
	if err != nil {
		log.Fatal("example ErrGroup finished with error: ", err.Error())
	}

	log.Println("example ErrGroup finished sucessfully")
}
