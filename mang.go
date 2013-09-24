// * 
// * Copyright 2013, Scott Cagno. All rights Reserved
// * License: sites.google.com/site/bsdc3license
// * 
// * -------
// * mang.go ::: database collection store manager
// * -------
// * 

import (
	"runtime"
	"bytes"
	"sync"
	"time"
)

package db

// store manager
type Manager struct {
	Stores 	map[string]*Store
	Marked 	map[string]int64
	mu 		sync.Mutex
}

// return new manager instance
func InitManager() *Manager {
	self := &Manager{
		Stores: make(map[string]*Store),
		Marked: make(map[string]int64),
	}
	go self.runGC(GC_RATE)
	return self
}

// run garbage collector
func (self *Manager) runGC(rate int64) {
	if len(self.Marked) > 0 {
		self.GC()
	}
	time.AfterFunc(time.Duration(rate)*time.Second, func(){ self.runGC(rate) })
}

// garbage collection
func (self *Manager) GC() {
	self.mu.Lock()
	for k, ttl := range self.Marked {
		if ttl <= time.Now().Unix() {
			delete(self.Stores, k)
			delete(self.Marked, k)
		}
	}
	self.mu.Unlock()
}

// check to see if a given key exists in stores
func (self *Manager) Has(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// add, or append a new store entry (does not overwrite)
func (self *Manager) Add(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// get an stores key list by key
func (self *Manager) Get(k string) [][]byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// delete an entire store
func (self *Manager) Del(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// set a store to expire in ttl seconds
func (self *Manager) Exp(k string, ttl int64) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// check the time to live for a store by key
func (self *Manager) Ttl(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}