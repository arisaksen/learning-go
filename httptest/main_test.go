package httptest

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestHandleGetUser(t *testing.T) {
	server := NewServer()
	testServer := httptest.NewServer(http.HandlerFunc(server.handleGetUser))
	numberOfRequests := 1000

	wg := &sync.WaitGroup{}
	for i := 0; i < numberOfRequests; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			id := i%100 + 1
			url := fmt.Sprintf("%s/?id=%d", testServer.URL, id)
			resp, err := http.Get(url)
			if err != nil {
				t.Error(err)
			}

			user := &User{}
			if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
				t.Error(err)
			}

			slog.Info(url, slog.Any("user", user))
		}(i)

		// to mock get request time
		time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()

	expectedNumberOfTimesToHitDb := 100
	if server.dbhit != expectedNumberOfTimesToHitDb {
		t.Errorf("Test FAIL: expected db hit: %d, actual: %d", expectedNumberOfTimesToHitDb, server.dbhit)
	}
}
