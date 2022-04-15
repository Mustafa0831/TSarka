package server

import (
	"os"
	"testing"
)

func TestServer(m *testing.M) {
	os.Exit(m.Run())
}
