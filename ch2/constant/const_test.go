package constant

import "testing"
const (
	One = 1
	Two = 2
	Three = 5
	Four = 4
)
func TestConst(t *testing.T) {
	t.Log(One,Two,Three,Four)
}