package services

import (
	"sync"
)

type NotificationService struct {
	mu      sync.Mutex
	clients map[chan string]bool
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		clients: make(map[chan string]bool),
	}
}

func (s *NotificationService) AddClient(c chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[c] = true
}

func (s *NotificationService) RemoveClient(c chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, c)
	close(c)
}

func (s *NotificationService) NotifyAll(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for client := range s.clients {
		client <- message
	}
}
