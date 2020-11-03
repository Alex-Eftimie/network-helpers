package main

import (
	"io"
	"sync"
)

// PipeStreams bidirectionally pipes streams
func PipeStreams(a, b io.ReadWriteCloser) {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer a.Close()
		defer b.Close()

		if _, err := io.Copy(a, b); err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		defer a.Close()
		defer b.Close()

		if _, err := io.Copy(b, a); err != nil {
			return
		}
	}()
	wg.Wait()
}
