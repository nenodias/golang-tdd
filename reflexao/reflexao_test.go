package reflexao

import (
	"reflect"
	"testing"
)

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

func verifyContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v contains '%s', but wasn't", haystack, needle)
	}
}

func TestPercorre(t *testing.T) {
	casos := []struct {
		Nome              string
		Entrada           interface{}
		ChamadasEsperadas []string
	}{
		{
			"Struct with a string field",
			struct {
				Nome string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Nome   string
				Cidade string
			}{"Chris", "Londres"},
			[]string{"Chris", "Londres"},
		},
		{
			"Struct with an int field",
			struct {
				Nome  string
				Idade int
			}{"Chris", 36},
			[]string{"Chris"},
		},
		{
			"Struct with embedded fields",
			Pessoa{"Chris", Perfil{33, "Londres"}},
			[]string{"Chris", "Londres"},
		},
		{
			"Pointers to things",
			&Pessoa{"Chris", Perfil{33, "Londres"}},
			[]string{"Chris", "Londres"},
		},
		{
			"Slices",
			[]Perfil{
				{33, "Londres"},
				{34, "São Paulo"},
			},
			[]string{"Londres", "São Paulo"},
		},
		{
			"Arrays",
			[2]Perfil{
				{33, "Londres"},
				{34, "São Paulo"},
			},
			[]string{"Londres", "São Paulo"},
		},
	}
	for _, teste := range casos {
		t.Run(teste.Nome, func(t *testing.T) {
			var got []string
			percorre(teste.Entrada, func(entrada string) {
				got = append(got, entrada)
			})

			if !reflect.DeepEqual(got, teste.ChamadasEsperadas) {
				t.Errorf("got '%v', want '%v'", got, teste.ChamadasEsperadas)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		var got []string
		input := map[string]string{
			"Baz": "Boz",
			"Foo": "Bar",
		}
		want := []string{"Bar", "Boz"}
		percorre(input, func(entrada string) {
			got = append(got, entrada)
		})
		for _, found := range got {
			verifyContains(t, want, found)
		}
	})
}
