package treap

import (
	"fmt"
	"reflect"
	"testing"
)

func assertMappness(t *testing.T, treap *Treap, m map[string]Thing) {
	if !reflect.DeepEqual(treap.ToMap(), m) {
		t.Errorf("%v should be %v", treap, m)
	}
	if treap.Size() != len(m) {
		t.Errorf("%v.Size() should be %v", treap, len(m))
	}
}

func TestTreapEach(t *testing.T) {
	treap := new(Treap)
	m := make(map[string]Thing)
	for i := 0; i < 10; i++ {
		treap.Put([]byte(fmt.Sprint(i)), i)
		if val, exists := treap.Get([]byte(fmt.Sprint(i))); val != i || !exists {
			t.Errorf("insert of %v failed!", i)
		}
		m[fmt.Sprint(i)] = i
	}
	assertMappness(t, treap, m)
	var collector []string
	treap.Up([]byte("5"), []byte("8"), func(key []byte, value Thing) {
		collector = append(collector, string(key))
	})
	if !reflect.DeepEqual(collector, []string{"5", "6", "7"}) {
		t.Errorf("%v is bad", collector)
	}
	collector = nil
	treap.Down([]byte("6"), []byte("3"), func(key []byte, value Thing) {
		collector = append(collector, string(key))
	})
	if !reflect.DeepEqual(collector, []string{"6", "5", "4"}) {
		t.Errorf("%v is bad", collector)
	}
}

func TestTreapBasicOps(t *testing.T) {
	treap := new(Treap)
	m := make(map[string]Thing)
	assertMappness(t, treap, m)
	if val, existed := treap.Get([]byte("key")); val != nil || existed {
		t.Errorf("should not have existed")
	}
	if old, existed := treap.Del([]byte("key")); old != nil || existed {
		t.Errorf("should not have existed")
	}
	if old, existed := treap.Put([]byte("key"), "value"); old != nil || existed {
		t.Errorf("should not have existed")
	}
	if val, existed := treap.Get([]byte("key")); val != "value" || !existed {
		t.Errorf("should not have existed")
	}
	m["key"] = "value"
	assertMappness(t, treap, m)
	if old, existed := treap.Put([]byte("key"), "value2"); old != "value" || !existed {
		t.Errorf("should have existed")
	}
	if val, existed := treap.Get([]byte("key")); val != "value2" || !existed {
		t.Errorf("should have existed")
	}
	m["key"] = "value2"
	assertMappness(t, treap, m)
	if old, existed := treap.Del([]byte("key")); old != "value2" || !existed {
		t.Errorf("should have existed")
	}
	delete(m, "key")
	assertMappness(t, treap, m)
	if old, existed := treap.Del([]byte("key")); old != nil || existed {
		t.Errorf("should not have existed")
	}
}

func benchTreap(b *testing.B, n int) {
	b.StopTimer()
	var v [][]byte
	for i := 0; i < n; i++ {
		v = append(v, []byte(fmt.Sprint(i)))
	}
	b.StartTimer()
	for j := 0; j < b.N/n; j++ {
		m := new(Treap)
		for i := 0; i < n; i++ {
			k := v[i]
			m.Put(k, i)
			j, _ := m.Get(k)
			if j != i {
				b.Error("should be same value")
			}
		}
	}
}

func benchMap(b *testing.B, n int) {
	b.StopTimer()
	var v []string
	for i := 0; i < n; i++ {
		v = append(v, fmt.Sprint(i))
	}
	b.StartTimer()
	for j := 0; j < b.N/n; j++ {
		m := make(map[string]Thing)
		for i := 0; i < n; i++ {
			k := v[i]
			m[k] = i
			j, _ := m[k]
			if j != i {
				b.Error("should be same value")
			}
		}
	}
}

func BenchmarkTreap10(b *testing.B) {
	benchTreap(b, 10)
}
func BenchmarkTreap100(b *testing.B) {
	benchTreap(b, 100)
}
func BenchmarkTreap1000(b *testing.B) {
	benchTreap(b, 1000)
}
func BenchmarkTreap10000(b *testing.B) {
	benchTreap(b, 10000)
}
func BenchmarkTreap100000(b *testing.B) {
	benchTreap(b, 100000)
}
func BenchmarkTreap1000000(b *testing.B) {
	benchTreap(b, 1000000)
}

func BenchmarkMap10(b *testing.B) {
	benchMap(b, 10)
}
func BenchmarkMap100(b *testing.B) {
	benchMap(b, 100)
}
func BenchmarkMap1000(b *testing.B) {
	benchMap(b, 1000)
}
func BenchmarkMap10000(b *testing.B) {
	benchMap(b, 10000)
}
func BenchmarkMap100000(b *testing.B) {
	benchMap(b, 1000000)
}
func BenchmarkMap1000000(b *testing.B) {
	benchMap(b, 1000000)
}
