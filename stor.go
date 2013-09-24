// * 
// * Copyright 2013, Scott Cagno. All rights Reserved
// * License: sites.google.com/site/bsdc3license
// * 
// * -------
// * stor.go ::: databse collection store
// * -------
// * 

package db

import (
	"runtime"
	"bytes"
	"sync"
	"time"
)

// collection store
type Store struct {
	Items	map[string][][]byte
	Marked	map[string]int64
	pk 		int64
	mu 		sync.Mutex	
}

// return new store instance
func InitStore() *Store {
	self := &Store{
		Items: make(map[string][][]byte),
		Marked: make(map[string]int64),
	}
	go self.runGC(GC_RATE)
	return self
}

// run garbage collector
func (self *Store) runGC(rate int64) {
	if len(self.Marked) > 0 {
		self.GC()
	}
	time.AfterFunc(time.Duration(rate)*time.Second, func(){ self.runGC(rate) })
}

// garbage collection
func (self *Store) GC() {
	self.mu.Lock()
	for k, ttl := range self.Marked {
		if ttl <= time.Now().Unix() {
			delete(self.Items, k)
			delete(self.Marked, k)
		}
	}
	self.mu.Unlock()
}

// check to see if a given key exists in items
func (self *Store) Has(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// add, or append a new entry (does not overwrite)
func (self *Store) Add(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// overwrite/update an item (considered not safe)
func (self *Store) Set(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// get an items data by key
func (self *Store) Get(k string) [][]byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// delete an item
func (self *Store) Del(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()	
	return nil
}

// set and item to expire in ttl seconds
func (self *Store) Exp(k string, ttl int64) []byte {
	self.mu.Lock()
	self.mu.Unlock()	
	return nil
}

// check the time to live for an item by key
func (self *Store) Ttl(k string) []byte {
	self.mu.Lock()
	self.mu.Unlock()	
	return nil
}

// locate an item by a search term, or keyword
func (self *Store) Loc(b []byte) []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// increment stores pk, and return 
func (self *Store) NextPk() []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}

// reset stores pk
func (self *Store) ResetPK() []byte {
	self.mu.Lock()
	self.mu.Unlock()
	return nil
}