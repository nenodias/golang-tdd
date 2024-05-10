package injecao

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Chris")

	got := buffer.String()
	want := "Olá, Chris"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
