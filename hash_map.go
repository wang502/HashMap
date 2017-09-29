package HashMap

import (
	"fmt"
)

// Pair represents a key-value pair stored in hash map
type Pair struct {
	key     string
	value   interface{}
	deleted bool
}

func (p *Pair) String() string {
	return fmt.Sprintf("key: %s, value: %v, deleted: %t\n", p.key, p.value, p.deleted)
}

// HashMap represents a fix-sized hash map
type HashMap struct {
	data     []*Pair
	numItems int
	size     int
}

// NewHashMap initializes a fixed-size hash map of given size
func NewHashMap(size int) *HashMap {
	return &HashMap{
		data:     make([]*Pair, size),
		numItems: 0,
		size:     size,
	}
}

// Set stores the given key/value pair in the hash map.
// Returns a boolean value indicating success/failure of the operation
func (hashMap *HashMap) Set(key string, value interface{}) bool {
	// hash map is full, fail to set value
	if hashMap.numItems == hashMap.size {
		return false
	}

	index := hashMap.firstFreeIndex(key)
	if index != -1 {
		slot := hashMap.data[index]

		//the slot is either completely empty or has been marked as deleted
		if slot == nil || (slot != nil && slot.deleted) {
			hashMap.numItems++
		}
		hashMap.data[index] = &Pair{key, value, false}
		return true
	}
	return false
}

// Get return the value associated with the given key, or nil if no value is set
func (hashMap *HashMap) Get(key string) interface{} {
	index := hashMap.findIndex(key)
	if index == -1 {
		return nil
	}

	return hashMap.data[index].value
}

// Delete deletes the value associated with the given key
// returning the value on success or null if the key has no value
func (hashMap *HashMap) Delete(key string) interface{} {
	index := hashMap.findIndex(key)
	var res interface{}
	if index != -1 {
		res = hashMap.data[index].value
		hashMap.data[index].deleted = true
		hashMap.numItems--
	}
	return res
}

// Load returns a float value representing the load factor of the hash map
func (hashMap *HashMap) Load() float64 {
	return float64(hashMap.numItems / hashMap.size)
}

// ----------------------------------------------------------------------------------//
// 							      Helper functions								     //
// ----------------------------------------------------------------------------------//

// firstFreeIndex finds the next available position for the given key
// available means: 1) the slot is completely empty
//				    2) the slot has been marked 'deleted'
func (hashMap *HashMap) firstFreeIndex(key string) int {
	startIndex := hashMap.hash(key)
	nextIndex := startIndex
	var pair *Pair
	for hashMap.data[nextIndex] != nil {
		pair = hashMap.data[nextIndex]
		if pair.key == key || pair.deleted {
			return nextIndex
		}

		nextIndex = hashMap.nextHash(nextIndex)

		// has loop through all buckets in the hash map and came back to start index
		if nextIndex == startIndex {
			return -1
		}
	}
	return nextIndex
}

// findIndex finds the given key's location in the hash map
func (hashMap *HashMap) findIndex(key string) int {
	startIndex := hashMap.hash(key)
	nextIndex := startIndex
	var pair *Pair
	for hashMap.data[nextIndex] != nil {
		pair = hashMap.data[nextIndex]
		if pair.key == key && !pair.deleted {
			return nextIndex
		}

		nextIndex = hashMap.nextHash(nextIndex)

		// has loop through all buckets in the hash map and came back to start index
		if nextIndex == startIndex {
			return -1
		}
	}
	return -1
}

func (hashMap *HashMap) nextHash(prevHash int) int {
	return (prevHash + 1) % hashMap.size
}

func (hashMap *HashMap) hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return sum % hashMap.size
}

func (hashMap *HashMap) String() {
	res := ""
	for i := 0; i < len(hashMap.data); i++ {
		pair := hashMap.data[i]
		if pair == nil {
			res += fmt.Sprintf("key: nil, value: nil\n")
		} else {
			res += pair.String()
		}
	}
	fmt.Print(res)
}
