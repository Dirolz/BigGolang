package main

import (
	"fmt"
	"homework1/internal/pkg/storage"
)

func main() {
	s, err := storage.NewStorage()
	if err == nil {
		s.Set("key1", "value1")
		s.Set("key2", "1")
		fmt.Println(*s.Get("key1"))
		fmt.Println(*s.Get("key2"))
		fmt.Println(s.Get("1111"))
		fmt.Println(s.GetType("key1"))
		fmt.Println(s.GetType("key2"))
	}
}
