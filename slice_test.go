package gomini

import (
	"testing"
)

func BenchmarkUniq(b *testing.B) {
	values := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "b"}
	for i := 0; i < b.N; i++ {
		values = Uniq(values)
		if len(values) != 2 {
			b.Fatal("values len is incorrect")
		}
	}
}

func TestUniq(t *testing.T) {
	values := []string{"a", "a", "a", "b", "b", "b", "c", "c", "c"}
	values = Uniq(values)
	if len(values) != 3 {
		t.Fatal("values len is incorrect")
	}
}
