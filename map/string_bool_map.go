package gmap

import (
	"sync"
)

func NewStringBool() stringBool {
	sbMap := stringBool{}
	sbMap.Clear()
	return sbMap
}
type stringBool struct {
	sync.RWMutex
	m map[string]bool
}
func (self *stringBool) Get(key string) (value bool, has bool) {
	self.RLock()
	value, has = self.m[key]
	self.RUnlock()
	return
}
func (self *stringBool) Has(key string) (has bool) {
	self.RLock()
	_, has = self.m[key]
	self.RUnlock()
	return
}
func (self *stringBool) Set(key string, value bool) {
	self.Lock()
	self.m[key] = value
	self.Unlock()
}
func (self *stringBool) Remove(key string) {
	self.Lock()
	delete(self.m, key)
	self.Unlock()
}
func (self *stringBool) Clear () {
	self.Lock()
	self.m = map[string]bool{}
	self.Unlock()
}
func (self stringBool) Size() int {
	self.RLock()
	size := len(self.m)
	self.RUnlock()
	return size
}
func (self *stringBool) Value() map[string]bool {
	return self.m
}