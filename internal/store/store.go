package store

import (
	"sync"
	"time"
)

type myStore struct {
	// в мапе храним данные
	store map[string]any
	// в мапе храним все TTL значения ключей
	// то есть ключ - время, когда его нужно удалить
	ttl map[string]time.Time
	// для работы в многопоточном режиме используем мьюьекс
	locker sync.RWMutex
}

func GetStore() *myStore {
	return &myStore{
		make(map[string]any),
		make(map[string]time.Time),
		sync.RWMutex{},
	}
}

// метод для добавления ключа с заданным ttl в секундах
func (s *myStore) Set(key string, value any, seconds int) {
	s.locker.Lock()
	s.store[key] = value
	s.ttl[key] = time.Now().Add(time.Second * time.Duration(seconds))
	s.locker.Unlock()
}

// метод для получения ключа
func (s *myStore) Get(key string) (any, error) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	now := time.Now()
	if s.ttl[key].Before(now) {
		return nil, ErrNoKey
	}
	return s.store[key], nil
}

// метод для удаления ключа
func (s *myStore) Delete(key string) {
	s.locker.Lock()
	delete(s.store, key)
	delete(s.ttl, key)
	s.locker.Unlock()
}
