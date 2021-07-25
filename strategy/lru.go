package strategy

import "fmt"

type LRU struct {}

func (l *LRU) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}