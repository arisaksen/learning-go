package main

import (
	"bytes"
	"sync"
	"time"
)

type dataStore struct {
	buff         *bytes.Buffer
	writeCounter *int
	mutex        *sync.Mutex
}

func newDataStore() dataStore {
	counter := 0
	return dataStore{
		buff:         new(bytes.Buffer),
		writeCounter: &counter,
		mutex:        new(sync.Mutex),
	}
}

func (ds dataStore) writeToBufferUnsafe(text string) {
	// Simulate some work
	time.Sleep(1 * time.Second)

	ds.buff.WriteString(text)
	ds.buff.WriteString("\n")
	*ds.writeCounter++
}

func (ds dataStore) writeToBufferSafe(text string) {
	// Simulate some work
	time.Sleep(1 * time.Second)

	// Lock the mutex before writing to the shared buffer
	ds.mutex.Lock()
	defer ds.mutex.Unlock() // Unlock the mutex after our 'writeToBuffer' function is complete

	ds.buff.WriteString(text)
	ds.buff.WriteString("\n")
	*ds.writeCounter++
}

func (ds dataStore) printTest() {
	time.Sleep(1 * time.Nanosecond)
	*ds.writeCounter++
}
