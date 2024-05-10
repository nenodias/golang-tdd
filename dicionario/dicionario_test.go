package dicionario

import "testing"

func TestBuscar(t *testing.T) {
	d := Dicionario{"test": "palavra"}
	t.Run("palavra encontrada", func(t *testing.T) {
		want := "palavra"
		got, err := d.Busca("test")
		if err != nil {
			t.Fatal("não foi possível encontrar a palavra adicionada: ", err)
		}
		if got != want {
			t.Errorf("got '%s', want '%s", got, want)
		}
	})
	t.Run("palavra não encontrada", func(t *testing.T) {
		want := ""
		got, err := d.Busca("teste")
		if err != ErroPalavraNaoExistente {
			t.Fatal("erro diferente do esperado: ", err)
		}
		if got != want {
			t.Errorf("got '%s', want '%s", got, want)
		}
	})
}

func TestAdiciona(t *testing.T) {
	d := Dicionario{"test": "palavra"}
	t.Run("palavra inexistente", func(t *testing.T) {
		err := d.Adciona("teste", "isso é apenas um teste")

		if err != nil {
			t.Fatal("não foi possível encontrar a palavra adicionada: ", err)
		}

		want := "isso é apenas um teste"
		got, errBusca := d.Busca("teste")

		if errBusca != nil {
			t.Fatal("não foi possível encontrar a palavra adicionada: ", err)
		}
		if got != want {
			t.Errorf("got '%s', want '%s", got, want)
		}
	})
	t.Run("palavra existente", func(t *testing.T) {
		err := d.Adciona("test", "isso é apenas um teste")

		if err != ErroPalavraExistente {
			t.Fatal("não foi possível adicionar a palavra: ", err)
		}
	})
}
