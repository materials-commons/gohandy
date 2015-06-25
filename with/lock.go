package with

import "sync"

// ReadLock calls fn after acquiring a read lock.
func ReadLock(mutex *sync.RWMutex, fn func()) {
	defer mutex.RUnlock()
	mutex.RLock()
	fn()
}

// WriteLock calls fn after acquiring a write lock.
func WriteLock(mutex *sync.RWMutex, fn func()) {
	defer mutex.Unlock()
	mutex.Lock()
	fn()
}

// An RWLocker is an interface that acquires
// a lock before calling the given func.
type RWLocker interface {
	WithReadLock(fn func())
	WithWriteLock(fn func())
}

// RWLock implements the RWLocker interface. It takes a function
// that wraps the function to call. This allows the user
// to perform common book keeping or checks before calling
// the function passed to WithReadLock or WithWriteLock.
type RWLock struct {
	mutex *sync.RWMutex
	fn    func(fn func())
}

func NewRWLock(fn func()) *RWLock {
	fnToCall := fn
	if fnToCall == nil {
		fnToCall = func(fun func()) {
			fun()
		}
	}
	return &RWLock{
		mutex: &sync.RWMutex{},
		fn:    fnToCall,
	}
}

// WithReadLock acquires a read lock and then calls
// the RWLock function passing it the fn.
func (lock *RWLock) WithReadLock(fn func()) {
	defer lock.mutex.RUnlock()
	lock.mutex.RLock()
	lock.fn(fn)
}

// WithWriteLock acquires a write lock and then calls
// the RWLock function passing it the fn.
func (lock *RWLock) WithWriteLock(fn func()) {
	defer lock.mutex.Unlock()
	lock.mutex.Lock()
	lock.fn(fn)
}

// Use uses the passed in fn, rather than the fn that
// that the RWLock was originally set up with.
func (lock *RWLock) Use(fn func()) *RWLock {
	rwLockToUse := &RWLock{
		mutex: lock.mutex,
		fn:    fn,
	}
	return rwLockToUse
}
