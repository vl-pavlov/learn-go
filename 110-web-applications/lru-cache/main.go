package main

import (
	"container/list"
	"fmt"
)

// Item represents a key-value pair stored in the LRU cache.
type Item struct {
	Key   string
	Value interface{}
}

// LRU is a struct that represents a Least Recently Used cache.
type LRU struct {
	capacity int                      // The maximum number of items the cache can hold.
	items    map[string]*list.Element // A map to store the items in the cache.
	queue    *list.List               // A list to keep track of the order of access for the items.
}

// NewLru creates a new LRU cache with the given capacity.
func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

// Set adds a new item to the cache or updates the value of an existing item.
// It returns true if the item was added or updated successfully, and false otherwise.
func (c *LRU) Set(key string, value interface{}) bool {
	defer fmt.Println("Item", key, "added")

	// If the item already exists in the cache, update its value and move it to the front of the queue.
	if element, exists := c.items[key]; exists {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	// If the cache is at capacity, remove the least recently used item before adding a new one.
	if c.queue.Len() == c.capacity {
		c.purge()
	}

	// Create a new item and add it to the cache.
	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element

	return true
}

// purge removes the least recently used item from the cache.
func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
		fmt.Println("Item", item.Key, "deleted")
	}
}

func main() {
	// Create a new LRU cache with a capacity of 3.
	lru := NewLru(3)

	// Add items to the cache.
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	lru.Set("key3", "value3")
	lru.Set("key4", "value4")
}
