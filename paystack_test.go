package paystack_go

import "testing"

func TestNewClient(t *testing.T) {
	c := New("sk_test_123456")
	if c == nil {
		t.Fatal("Expected 'Client' instance, got nil")
	}
}
