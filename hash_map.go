package HashMap

// ListNode represents a node in a linked list
type ListNode struct {
	key   string
	value interface{}
	next  *ListNode
}

// HashMap represents a fix-sized hash map
type HashMap struct {
	data     []*ListNode
	numItems int
	size     int
}

// NewHashMap initializes a fixed-size hash map of given size
func NewHashMap(size int) *HashMap {
	return &HashMap{
		data:     make([]*ListNode, size),
		numItems: 0,
		size:     size,
	}
}

// Set stores the given key/value pair in the hash map.
// Returns a boolean value indicating success/failure of the operation
func (hashMap *HashMap) Set(key string, value interface{}) bool {
	if hashMap.numItems == hashMap.size {
		return false
	}

	hashKey := hashMap.hash(key)
	headNode := hashMap.data[hashKey]
	for headNode != nil {
		if headNode.key == key {
			headNode.value = value
			return true
		}
		headNode = headNode.next
	}

	newNode := &ListNode{key, value, nil}
	newNode.next = hashMap.data[hashKey]
	hashMap.data[hashKey] = newNode
	hashMap.numItems++
	return true
}

// Get return the value associated with the given key, or nil if no value is set
func (hashMap *HashMap) Get(key string) interface{} {
	hashKey := hashMap.hash(key)
	headNode := hashMap.data[hashKey]
	for headNode != nil {
		if headNode.key == key {
			return headNode.value
		}
		headNode = headNode.next
	}

	return nil
}

// Delete deletes the value associated with the given key
// returning the value on success or null if the key has no value
func (hashMap *HashMap) Delete(key string) interface{} {
	hashKey := hashMap.hash(key)
	curNode := hashMap.data[hashKey]
	var preNode *ListNode
	// traverse the list to find node with given key
	for curNode != nil {
		if curNode.key == key {
			if preNode != nil {
				preNode.next = curNode.next
			} else {
				hashMap.data[hashKey] = curNode.next
			}
			return curNode.value
		}
		preNode = curNode
		curNode = curNode.next
	}

	// when the key is not found
	return nil
}

// Load returns a float value representing the load factor of the hash map
func (hashMap *HashMap) Load() float64 {
	return float64(hashMap.numItems / hashMap.size)
}

func (hashMap *HashMap) hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return sum % hashMap.size
}
