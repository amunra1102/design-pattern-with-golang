package strategy

import "fmt"

type FIFO struct {}

func (f *FIFO) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}
