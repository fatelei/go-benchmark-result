package gobenchmarkresult

import (
	"sync"
	"testing"
)

func BenchmarkReadSyncMap(t *testing.B) {
	var sb sync.Map
	for i := 0; i < t.N; i++ {
		sb.Load("foo")
	}
}

func BenchmarkReadManualSyncMap(t *testing.B) {
	sb := make(map[string]string)
	var lock sync.RWMutex
	for i := 0; i < t.N; i++ {
		lock.RLock()
		_ = sb["foo"]
		lock.RUnlock()
	}
}
