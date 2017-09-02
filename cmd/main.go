package main

import (
	"fmt"

	"github.com/wang502/HashMap"
)

func main() {
	hashMap := HashMap.NewHashMap(5)
	hashMap.Set("name", "Seth")
	hashMap.Set("age", 23)
	fmt.Printf("age: %v\n", hashMap.Get("age"))
	fmt.Printf("name: %v\n", hashMap.Get("name"))

	hashMap.Delete("name")
	fmt.Printf("name: %v", hashMap.Get("name"))
}
