package gen

import "testing"

func BenchmarkGetCommands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetCommands(nil, nil)
	}
}
