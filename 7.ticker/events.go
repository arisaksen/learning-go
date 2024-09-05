package main

import (
	"time"
)

type Manager struct {
	counterId  int
	events     chan Event
	ticker     *time.Ticker
	tickerDone chan bool
}
type Event struct {
	id      int
	message string
	time    time.Time
}

func CreateManager(timeDelta time.Duration) Manager {
	newTicker := time.NewTicker(timeDelta)
	manager := Manager{
		events: make(chan Event),
		ticker: newTicker,
	}

	return manager
}

func (m *Manager) Stream() <-chan Event {

	go func() {
		for {
			select {
			case <-m.tickerDone:
				return
			case t := <-m.ticker.C:
				m.counterId++
				m.events <- Event{
					id:      m.counterId,
					message: "event",
					time:    t,
				}
			}
		}
	}()

	return m.events
}
