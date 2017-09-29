package main

import (
	"fmt"

	"github.com/wang502/HashMap"
)

// Example usage of HashMap library
func main() {
	hashMap := HashMap.NewHashMap(10)
	hashMap.Set("name", "Seth")
	hashMap.Set("age", 23)
	fmt.Printf("age: %v\n", hashMap.Get("age"))
	fmt.Printf("name: %v\n", hashMap.Get("name"))

	hashMap.Set("gea", 43)
	fmt.Printf("gea: %v\n", hashMap.Get("gea"))
	/*hashMap.Delete("name")
	fmt.Printf("name: %v\n", hashMap.Get("name"))
	*/
	hashMap.Set("name", "Seth")
	hashMap.String()
	fmt.Println("")

	hashMap.Delete("name")
	hashMap.String()
	fmt.Println("")

	fmt.Printf("name: %v\n", hashMap.Get("name"))
	fmt.Printf("gea: %v\n", hashMap.Get("gea"))
	fmt.Println("")

	hashMap.Set("amen", "master")
	hashMap.String()
	fmt.Println("")

	hashMap.Set("gea", 100)
	hashMap.String()
	fmt.Println("")

	hashMap.Set("aeg", 200)
	hashMap.String()
	fmt.Println("")

	fmt.Printf("aeg: %v\n", hashMap.Get("aeg"))

	hashMap.Set("Aegon", "dragon")
	hashMap.String()
	fmt.Println("")

	hashMap.Delete("aeg")
	fmt.Println(hashMap.Set("Aegon", "dragon"))
	hashMap.String()
	fmt.Println("")

	fmt.Printf("aeg: %v\n", hashMap.Get("aeg"))
	hashMap.Set("age", 300)
	hashMap.String()
	fmt.Println("")

	fmt.Printf("amen: %v\n", hashMap.Get("amen"))
	hashMap.Set("amen", "100")
	fmt.Printf("amen: %v\n", hashMap.Get("amen"))

	hashMap.Delete("amen")

	hashMap.Set("mean", "char")
	hashMap.String()
	fmt.Println("")

	hashMap.Set("aeg", "400")
	hashMap.String()
	fmt.Println("")
}
