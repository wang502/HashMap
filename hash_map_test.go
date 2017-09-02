package HashMap

import (
	"testing"
)

func TestSet(t *testing.T) {
	hashMap := NewHashMap(5)
	hashMap.Set("KPCB", "Pinterest")
	hashMap.Set("Menlo Park", "KPCB")

	v1 := hashMap.Get("KPCB")
	v2 := hashMap.Get("Menlo Park")
	if v1 != "Pinterest" {
		t.Errorf("key %s's value should be Pinterest, but %s is returned", "KPCB", v1)
	}
	if v2 != "KPCB" {
		t.Errorf("key %s's value should be KPCB, but %s is returned", "Menlo Park", v1)
	}
}

func TestGet(t *testing.T) {
	hashMap := NewHashMap(5)
	v1 := hashMap.Get("KPCB")
	if v1 != nil {
		t.Errorf("key %s is not in hash map, but value %s is returned", "KPCB", v1)
	}

	hashMap.Set("KPCB", "Pinterest")
	v2 := hashMap.Get("KPCB")
	if v2 != "Pinterest" {
		t.Errorf("key %s's value should be Pinterest, but %s is returned", "KPCB", v2)
	}
}

func TestDelete(t *testing.T) {
	hashMap := NewHashMap(5)
	hashMap.Set("KPCB", "Pinterest")
	hashMap.Set("Menlo Park", "KPCB")
	v1 := hashMap.Get("KPCB")
	v2 := hashMap.Get("Menlo Park")
	if v1 != "Pinterest" {
		t.Errorf("key %s's value should be Pinterest, but %s is returned", "KPCB", v1)
	}
	if v2 != "KPCB" {
		t.Errorf("key %s's value should be KPCB, but %s is returned", "Menlo Park", v1)
	}

	deletedV1 := hashMap.Delete("KPCB")
	if deletedV1 != "Pinterest" {
		t.Errorf("key %s's value Pinterest is deleted from hash map, but %s is returned", "KPCB", deletedV1)
	}

	v3 := hashMap.Get("KPCB")
	if v3 != nil {
		t.Errorf("key %s's value should be nil, but %s is returned", "KPCB", v3)
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
