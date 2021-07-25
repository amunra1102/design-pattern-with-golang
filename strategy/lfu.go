package strategy

import "fmt"

type LFU struct {}

func (l *LFU) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
}
