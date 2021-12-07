package manager

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/mediator/code/interfaces"
	"sync"
)

type stationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []interfaces.Train
}

func NewStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *stationManager) CanArrival(t interfaces.Train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) NotifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
