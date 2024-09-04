package main

import (
	"fmt"
	"sync"
)

// Mutex - Concurrency Safety: By ensuring the mutex is properly locked/unlocked, you maintain safe and correct concurrency.
// If we were to remove the lock from the 'writeToBuffer' function. We will see that the length of the buffer will vary.
// Might have to run multiple times
func main() {
	myDataStoreMutex := newDataStore()
	myDataStoreNoLock := newDataStore()
	var wg sync.WaitGroup
	texts := []string{
		"Hello, World!",
		"Goroutines are fun.",
		"Concurrency is powerful.",
		"Go is great!",
	}

	// Launch multiple goroutines
	for _, text := range texts {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go myDataStoreNoLock.writeToBufferUnsafe(text)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			myDataStoreMutex.writeToBufferSafe(text)
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("NO LOCK- - - - - - - - - - - - - - - - - -")
	fmt.Println(myDataStoreNoLock.buff.String())
	fmt.Println(myDataStoreNoLock.buff.Len())
	fmt.Println(*myDataStoreNoLock.writeCounter)

	fmt.Println("MUTEX - - - - - - - - - - - - - - - - - -")
	fmt.Println(myDataStoreMutex.buff.String())
	fmt.Println(myDataStoreMutex.buff.Len())
	fmt.Println(*myDataStoreMutex.writeCounter)
}
