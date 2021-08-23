package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	gerrors := make(chan error)
	wgDone := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		wg.Done()
	}()
	go func() {
		err := returnError()
		if err != nil {
			gerrors <- err
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(wgDone)
	}()
	select {
	case <-wgDone:
		break
	case err := <-gerrors:
		close(gerrors)
		fmt.Println(err)
	}
}

func returnError() error {
	return errors.New("this is a new error")
}
