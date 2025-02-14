package main

import (
	"go.uber.org/goleak"
	"testing"
)

func TestA(t *testing.T) {
	defer goleak.VerifyNone(t)
	Sos()

}
