package hello

import (
	"testing"
)

func TestHelloGrpc(t *testing.T) {
	msg := HelloGrpc()

	if msg != "Hello, gRPC!" {
		t.Fatalf(`HelloGrpc() returned "%s", but should return "Hello, gRPC!" `, msg)
	}
}
