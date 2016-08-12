package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func main() {
	exampleWaitgroup()
	//exampleErrgroup()
	//exampleErrgroupTimeout()
	//exampleErrgroupCancel()
}

//
// Example Waitgroup
//

func work1(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	wg.Done()
}

func exampleWaitgroup() {
	var wg sync.WaitGroup

	log.Println("example WaitGroup started")

	wg.Add(2)
	go work1(&wg)
	go work1(&wg)

	wg.Wait()
	log.Println("example WaitGroup finished")
}

//
// Example errgroup
//

func work2a() error {
	time.Sleep(1 * time.Second)
	return errors.New("we got an error")
}

func work2b() error {
	time.Sleep(2 * time.Second)
	return nil
}

func exampleErrgroup() {
	var eg errgroup.Group

	log.Println("example ErrGroup started")

	eg.Go(work2a)
	eg.Go(work2b)

	err := eg.Wait()
	if err != nil {
		log.Fatal("finished with error: ", err.Error())
	}

	log.Println("finished sucessfully")
}

//
// Example errgroup with context: timeout
//

func work3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func exampleErrgroupTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	log.Println("example errgroup with timeout")

	eg.Go(func() error {
		work3(ctx)
		return ctx.Err()
	})

	err := eg.Wait()
	if err != nil {
		log.Fatal("finished with error: ", err.Error())
	}

	log.Println("finished sucessfully")
}

//
// Example errgroup with context: cancel
//

func exampleErrgroupCancel() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	log.Println("example errgroup with cancel")

	eg.Go(func() error {
		work3(ctx)
		return ctx.Err()
	})
	eg.Go(work2a)

	err := eg.Wait()
	if err != nil {
		log.Fatal("finished with error: ", err.Error())
	}

	log.Println("finished sucessfully")
}
