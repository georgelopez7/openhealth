package test

import (
	"sync"
	"testing"

	"openhealth/internal/event"

	"github.com/stretchr/testify/require"
)

func TestEventMutex_Add(t *testing.T) {
	m := event.NewEventMutex()

	require.True(t, m.Add("event-1"))
	require.False(t, m.Add("event-1"))
	require.True(t, m.Add("event-2"))
}

func TestEventMutex_Remove(t *testing.T) {
	m := event.NewEventMutex()

	require.True(t, m.Add("event-1"))
	m.Remove("event-1")
	require.True(t, m.Add("event-1"))
}

func TestEventMutex_InProgress(t *testing.T) {
	m := event.NewEventMutex()

	require.False(t, m.InProgress("event-1"))

	m.Add("event-1")
	require.True(t, m.InProgress("event-1"))

	m.Remove("event-1")
	require.False(t, m.InProgress("event-1"))
}

func TestEventMutex_ConcurrentAdd(t *testing.T) {
	m := event.NewEventMutex()

	require.True(t, m.Add("event-1"))

	var wg sync.WaitGroup
	attempted := make(chan struct{}, 10)

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !m.Add("event-1") {
				attempted <- struct{}{}
			}
		}()
	}

	// Wait for all goroutines to attempt adding the event.
	for range 10 {
		<-attempted
	}

	m.Remove("event-1")
	wg.Wait()
	close(attempted)

	require.False(t, m.InProgress("event-1"))
}
