package other

type LRUCache struct {
	hash     map[int]*LRUNode
	head     *LRUNode
	last     *LRUNode
	length   int
	capacity int
}

type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hash:     map[int]*LRUNode{},
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	i := this.get(key)
	if i != nil {
		return i.value
	}

	return -1
}

func (this *LRUCache) get(key int) *LRUNode {
	i, ok := this.hash[key]
	if !ok {
		return nil
	}

	if this.length == 1 || this.head == i {
		return i
	}

	prev := i.prev
	next := i.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	if this.head != i {
		i.prev = nil
		i.next = this.head
		this.head.prev = i
		this.head = i
	}

	if this.last == i && prev != nil {
		this.last = prev
		this.last.next = nil
	}

	return i
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	i := this.get(key)

	if i != nil {
		i.value = value
		return
	}

	i = &LRUNode{
		key:   key,
		value: value,
	}
	this.hash[key] = i
	this.length++

	if this.length == 1 {
		this.head = i
		this.last = i
		return
	}

	this.head.prev = i
	i.next = this.head
	this.head = i

	if this.length > this.capacity {
		delete(this.hash, this.last.key)
		this.last = this.last.prev
		if this.last != nil {
			this.last.next = nil
		}
		this.length--
	}
}
