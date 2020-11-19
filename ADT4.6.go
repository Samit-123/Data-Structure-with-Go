package main

import "fmt"

type Counter map[interface{}]int

func (s *Counter) Add(key interface{}) {
	(*s)[key] += 1
}
func (s *Counter) Find(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}
func (s *Counter) Get(key interface{}) (int, bool) {
	value, ok := (*s)[key]
	return value, ok
}
func main() {
	c := make(Counter)
	c.Add("a")
	c.Add("b")
	c.Add("a")
	fmt.Println(c.Get("a"))
	fmt.Println(c.Get("b"))
	fmt.Println(c.Get("c"))
}
