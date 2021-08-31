package main

import "sync"

// FooMap 代表互斥锁
type FooMap struct {
	sync.Mutex
	data map[int]int
}

// BarRwMap 代表读写锁
type BarRwMap struct {
	sync.RWMutex
	data map[int]int
}

var fooMap *FooMap
var barRwMap *BarRwMap
var syncMap *sync.Map

func init() {
	fooMap = &FooMap{data: make(map[int]int, 100)}
	barRwMap = &BarRwMap{data: make(map[int]int, 100)}
	syncMap = &sync.Map{}
}

// buildinRwMapStore 模拟读写锁写入
func buildinRwMapStore(k, v int) {
	barRwMap.Lock()
	defer barRwMap.Unlock()
	barRwMap.data[k] = v
}

// buildinRwMapLookup 模拟读写锁读取
func buildinRwMapLookup(k int) int {
	barRwMap.RLock()
	defer barRwMap.RUnlock()
	if v, ok := barRwMap.data[k]; !ok {
		return -1
	} else {
		return v
	}
}

// buildinRwMapDelete 模拟读写锁删除
func buildinRwMapDelete(k int) {
	barRwMap.Lock()
	defer barRwMap.Unlock()
	if _, ok := barRwMap.data[k]; !ok {
		return
	}
	delete(barRwMap.data, k)
}

func buildinMapStore(k, v int) {
	fooMap.Lock()
	defer fooMap.Unlock()
	fooMap.data[k] = v
}

func buildinMapLookUp(k int) int {
	fooMap.Lock()
	defer fooMap.Unlock()
	if v, ok := fooMap.data[k]; !ok {
		return -1
	} else {
		return v
	}
}

func buildinMapDelete(k int) {
	fooMap.Lock()
	defer fooMap.Unlock()
	if _, ok := fooMap.data[k]; !ok {
		return
	}
	delete(fooMap.data, k)
}

func buildinSyncMapStore(k, v int) {
	syncMap.Store(k, v)
}

func buildinSyncMapLookUp(k int) int {
	t, ok := syncMap.Load(k)
	if ok {
		return t.(int)
	}
	return -1
}

func buildinSyncMapDelete(k int) {
	syncMap.Delete(k)
}

func main() {

}
