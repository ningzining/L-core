package server

import (
	"testing"
)

func TestStart(t *testing.T) {
	New(8080).Start()
}
