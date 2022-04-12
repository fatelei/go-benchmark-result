package gobenchmarkresult

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringBuilder(t *testing.B) {
	for i := 0; i < t.N; i++ {
		var sb strings.Builder
		for j := 0; j < 10000; j++ {
			sb.WriteString("哈")
		}
	}
}

func BenchmarkByteBufferBuilder(t *testing.B) {
	for i := 0; i < t.N; i++ {
		var sb bytes.Buffer
		for j := 0; j < 10000; j++ {
			sb.WriteString("哈")
		}
	}
}
