package HashMap

import (
	"testing"
)

var (
	size = 50
)

func TestSet(t *testing.T) {
	hashMap := NewHashMap(size)
	hashMap.Set("KPCB", "Pinterest")
	hashMap.Set("year", 2017)

	v1 := hashMap.Get("KPCB")
	v2 := hashMap.Get("year")
	if v1 != "Pinterest" {
		t.Errorf("key %s's value should be Pinterest, but %v is returned", "KPCB", v1)
	}
	if v2 != 2017 {
		t.Errorf("key %s's value should be 2017, but %v is returned", "year", v2)
	}

	hashMap.Set("year", 2018)
	v3 := hashMap.Get("year")
	if v3 != 2018 {
		t.Errorf("key %s's value should be 2018, but %v is returned", "year", v2)
	}
}

func TestGet(t *testing.T) {
	hashMap := NewHashMap(size)
	v1 := hashMap.Get("KPCB")
	if v1 != nil {
		t.Errorf("key %s is not in hash map, but value %v is returned", "KPCB", v1)
	}

	hashMap.Set("KPCB", "Pinterest")
	v2 := hashMap.Get("KPCB")
	if v2 != "Pinterest" {
		t.Errorf("key %s's value should be Pinterest, but %v is returned", "KPCB", v2)
	}
}

func TestDelete(t *testing.T) {
	hashMap := NewHashMap(size)
	hashMap.Set("KPCB", "Pinterest")
	hashMap.Set("year", 2017)
	v1 := hashMap.Get("KPCB")
	v2 := hashMap.Get("year")
	if v1 != "Pinterest" {
		t.Errorf("key %s's value should be Pinterest, but %v is returned", "KPCB", v1)
	}
	if v2 != 2017 {
		t.Errorf("key %s's value should be 2017, but %v is returned", "year", v2)
	}

	deletedV1 := hashMap.Delete("KPCB")
	if deletedV1 != "Pinterest" {
		t.Errorf("key %s's value Pinterest is deleted from hash map, but %v is returned", "KPCB", deletedV1)
	}

	v3 := hashMap.Get("KPCB")
	if v3 != nil {
		t.Errorf("key %s's value should be nil, but %v is returned", "KPCB", v3)
	}

	if hashMap.numItems != 1 {
		t.Errorf("number of items should be %d, but %d is returned", 1, hashMap.numItems)
	}
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	hashMap := NewHashMap(10)
	hashMap.Set("KPCB", "Pinterest")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hashMap.Get("KPCB")
	}
}

func BenchmarkSet(b *testing.B) {
	b.StopTimer()
	hashMap := NewHashMap(10)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hashMap.Set("KPCB", "Pinterest")
	}
}

func BenchmarkSetDelete(b *testing.B) {
	b.StopTimer()
	hashMap := NewHashMap(10)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hashMap.Set("KPCB", "Pinterest")
		hashMap.Delete("KPCB")
	}
}

func BenchmarkBuiltinMapGet(b *testing.B) {
	b.StopTimer()
	hashMap := make(map[string]interface{})
	hashMap["KPCB"] = "Pinterest"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = hashMap["KPCB"]
	}
}
