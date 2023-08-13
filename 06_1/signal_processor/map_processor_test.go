package signalprocessor

import "testing"

func TestUpdateWindowArraForMap(t *testing.T) {
	processor := MapProcessor{}
	processor.Initialise(4, byte('z'-'a')+1, byte('a'))
	// Feed it some values until we start looping.
	processor.updateWindowArray('a')
	processor.updateWindowArray('b')
	processor.updateWindowArray('c')
	processor.updateWindowArray('d')
	result := processor.updateWindowArray('e')
	if result != 'a' {
		t.Fatalf("Expected updateWindowArray to return a, got %c", result)
	}
	result = processor.updateWindowArray('e')
	if result != 'b' {
		t.Fatalf("Expected updateWindowArray to return b, got %c", result)
	}
	result = processor.updateWindowArray('e')
	if result != 'c' {
		t.Fatalf("Expected updateWindowArray to return c, got %c", result)
	}
	result = processor.updateWindowArray('e')
	if result != 'd' {
		t.Fatalf("Expected updateWindowArray to return d, got %c", result)
	}
	result = processor.updateWindowArray('e')
	if result != 'e' {
		t.Fatalf("Expected updateWindowArray to return e for the first time, got %c", result)
	}
	result = processor.updateWindowArray('e')
	if result != 'e' {
		t.Fatalf("Expected updateWindowArray to return e for the second time, got %c", result)
	}
}
