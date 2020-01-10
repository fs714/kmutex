package kmutex

import (
	"sync"
)

type keyLocker struct {
	count      int
	innerMutex sync.Mutex
}

type KMutex struct {
	outMutex   sync.Mutex
	kLockerMap map[string]*keyLocker
}

func New() *KMutex {
	return &KMutex{
		kLockerMap: make(map[string]*keyLocker),
	}
}

func (km *KMutex) Lock(key string) {
	km.outMutex.Lock()
	kl, ok := km.kLockerMap[key]
	if !ok {
		kl = new(keyLocker)
		km.kLockerMap[key] = kl
	}
	kl.count++
	km.outMutex.Unlock()

	kl.innerMutex.Lock()
}

func (km *KMutex) Unlock(key string) {
	km.outMutex.Lock()
	s, ok := km.kLockerMap[key]
	if !ok {
		km.outMutex.Unlock()
		return
	}
	s.count--
	if s.count == 0 {
		delete(km.kLockerMap, key)
	}
	km.outMutex.Unlock()

	s.innerMutex.Unlock()
}
