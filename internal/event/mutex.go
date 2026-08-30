package event

import "sync"

type EventMutex struct {
	mu     sync.RWMutex
	events map[string]bool
}

func NewEventMutex() *EventMutex {
	return &EventMutex{
		events: make(map[string]bool),
	}
}

// Add - marks the event as being processed.
func (m *EventMutex) Add(eventID string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.events[eventID] {
		return false
	}

	m.events[eventID] = true
	return true
}

// Remove - removes the event from the mutex
func (m *EventMutex) Remove(eventID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.events, eventID)
}

// InProgress - checks if the event is currently being processed
func (m *EventMutex) InProgress(eventID string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.events[eventID]
}
